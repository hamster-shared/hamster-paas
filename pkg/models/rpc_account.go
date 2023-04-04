package models

import (
	"fmt"
	"hamster-paas/pkg/application"

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
	var chains []RpcChain
	err = a.db.Model(&RpcChain{}).Find(&chains).Error
	if err != nil {
		return err
	}
	for _, chain := range chains {
		name := fmt.Sprintf("%s:%s", chain.Name, chain.Network)
		_, err = a.CreateAppByString(name, "", chain.Name, chain.Network)
		if err != nil {
			return err
		}
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

func (a *RpcAccount) GetApp(id int) (*RpcApp, error) {
	return getApp(a.Address, id)
}

func (a *RpcAccount) GetAppByName(name string) (*ApiResponseRpcApp, error) {
	return getAppByName(a.Address, name)
}

func (a *RpcAccount) GetAppByChainNetwork(chain ChainType, network NetworkType) (*ApiResponseRpcApp, error) {
	return getAppByChainNetwork(a.Address, chain, network)
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
