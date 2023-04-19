package models

import (
	"fmt"
	"hamster-paas/pkg/application"
	"strings"
	"time"

	"gorm.io/gorm"
)

type RpcAccount struct {
	gorm.Model
	db      *gorm.DB
	Address string `json:"address"`
}

func NewRpcAccount(address string) (*RpcAccount, error) {
	a := &RpcAccount{
		Address: address,
	}
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %s", err)
	}
	a.db = db
	return a, a.save()
}

func (a *RpcAccount) save() error {
	err := a.db.Create(a).Error
	if err != nil {
		return err
	}
	return nil
}

func GetRpcAccount(address string) (*RpcAccount, error) {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %s", err)
	}
	var account RpcAccount
	if err := db.Where("address = ?", address).First(&account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return NewRpcAccount(address)
		}
		return nil, fmt.Errorf("failed to get account: %s", err)
	}
	account.db = db
	return &account, nil
}

func (a *RpcAccount) CreateAppByString(name, description, chain, network string) (*RpcApp, error) {
	chainType, err := ParseChainType(chain)
	if err != nil {
		return nil, err
	}
	networkType, err := ParseNetworkType(network)
	if err != nil {
		return nil, err
	}
	return a.CreateApp(name, description, chainType, networkType)
}

func (a *RpcAccount) CreateApp(name, description string, chain ChainType, network NetworkType) (*RpcApp, error) {
	return newApp(a.Address, name, description, chain, network)
}

func (a *RpcAccount) DeleteApp(id int) error {
	return deleteApp(a.Address, id)
}

func (a *RpcAccount) GetApps() ([]*ApiResponseRpcApp, error) {
	return getApps(a.Address)
}

func (a *RpcAccount) GetOverview(network string) (*ApiResponseOverview, error) {
	var apps []*ApiResponseRpcApp
	var err error
	if strings.ToLower(network) == "mainnet" {
		apps, err = filterAppsMainnet(a.Address)
	} else if strings.ToLower(network) == "testnet" {
		apps, err = filterAppsTestnet(a.Address)
	} else {
		return nil, fmt.Errorf("invalid network type: %s, only 'mainnet' or 'testnet'", network)
	}
	if err != nil {
		return nil, err
	}
	var apiResponseOverview ApiResponseOverview
	apiResponseOverview.Network = network
	// 首先把非当前网络的过滤掉
	for _, v := range apps {
		// 过滤出种类
		if !contains(apiResponseOverview.LegendData, v.Name) {
			apiResponseOverview.LegendData = append(apiResponseOverview.LegendData, v.Name)
		}
	}
	for i := 0; i < 7; i++ {
		// 获取 utc+0 的时间
		nowUnix := time.Now().Unix()
		yesterdayStartUnix := nowUnix - int64(nowUnix%86400) - 86400 - int64(i*86400)
		apiResponseOverview.XaxisData = append(apiResponseOverview.XaxisData, time.Unix(yesterdayStartUnix, 0).UTC().Format(time.RFC3339))
	}
	apiResponseOverview.XaxisData = reverseString(apiResponseOverview.XaxisData)

	// 根据时间过滤，只保留最近 7 天的请求事件，每天一个，最后成一个数组，最终只要出现的次数而已
	for _, v := range apiResponseOverview.LegendData {
		var serie Serie
		serie.Name = v
		for _, date := range apiResponseOverview.XaxisData {
			var count int
			for _, app := range apps {
				if app.Name == v {
					for _, d := range app.DaylyRequests7Days {
						if strings.Contains(d.StartTime, date) {
							count += int(d.RequestNumber)
						}
					}
				}
			}
			serie.Data = append(serie.Data, count)
		}
		apiResponseOverview.SeriesData = append(apiResponseOverview.SeriesData, serie)
	}
	return &apiResponseOverview, nil
}

func (a *RpcAccount) GetAppsWithPagination(p *Pagination) ([]*ApiResponseRpcApp, *Pagination, error) {
	return getAppsPagination(a.Address, p)
}

func (a *RpcAccount) GetApp(id int) (*RpcApp, error) {
	return getApp(a.Address, id)
}

func (a *RpcAccount) GetAppByName(name string) (*ApiResponseRpcApp, error) {
	return getAppByName(a.Address, name)
}

func (a *RpcAccount) GetAppByChainNetwork(chain ChainType, network NetworkType) (*ApiResponseRpcApp, error) {
	return getAppByChainNetwork(a.Address, chain, network)
}

func (a *RpcAccount) GetAppBaseInfoByChainNetwork(chain ChainType, network NetworkType) (*ApiResponseRpcApp, error) {
	return getAppBaseInfoByChainNetwork(a.Address, chain, network)
}

func (a *RpcAccount) GetAppRequestLogs(appKey string, p Pagination) ([]*RpcAppRequestLog, *Pagination, error) {
	if !accountHaveApp(a.Address, appKey) {
		return nil, &p, fmt.Errorf("account have not app %s", appKey)
	}
	rpcApp := &RpcApp{
		ApiKey: appKey,
	}
	return rpcApp.getAppRequestLogs(appKey, p)
}

// 查看某个键是否在列表里
func contains(list []string, key string) bool {
	for _, v := range list {
		if v == key {
			return true
		}
	}
	return false
}

func reverseString(in []string) []string {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
	return in
}
