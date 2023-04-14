package models

import (
	"crypto/md5"
	"fmt"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/utils/logger"
	"time"

	"github.com/google/uuid"
	"github.com/meilisearch/meilisearch-go"
	"gorm.io/gorm"
)

type RpcApp struct {
	AppID         int64     `json:"app_id"`
	Account       string    `json:"account,omitempty"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Chain         string    `json:"chain"`
	Network       string    `json:"network"`
	ApiKey        string    `json:"api_key"`
	HttpLink      string    `json:"http_link"`
	WebsocketLink string    `json:"websocket_link"`
	CreatedAt     time.Time `json:"created_at"`
}

type ApiResponseRpcApp struct {
	*RpcApp
	CodeExamples       RpcCodeExample `json:"code_examples,omitempty"`
	TotalRequests24h   int64          `json:"total_requests_24h,omitempty"`
	DaylyRequests7Days []DateRequest  `json:"dayly_requests_7days,omitempty"`
	TotalRequestsAll   int64          `json:"total_requests_all,omitempty"`
}

type DateRequest struct {
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	RequestNumber int64  `json:"request"`
}

type RpcCodeExample struct {
	Js     string `json:"js"`
	Cli    string `json:"cli"`
	Go     string `json:"go"`
	Python string `json:"python"`
}

func newCodeExample() RpcCodeExample {
	return RpcCodeExample{
		Js:     "this is js code example",
		Cli:    "this is cli code example",
		Go:     "this is go code example",
		Python: "this is python code example",
	}
}

func newApp(account string, name, description string, chain ChainType, network NetworkType) (*RpcApp, error) {
	a := &RpcApp{
		Account:     account,
		Name:        name,
		Description: description,
		Chain:       chain.String(),
		Network:     network.String(),
	}
	err := a.generateKey()
	if err != nil {
		return nil, err
	}
	httpLink, wsLink, err := GetChainLink(chain, network)
	if err != nil {
		logger.Errorf("failed to get chain link: %s", err)
	}
	if httpLink != "" {
		a.HttpLink = fmt.Sprintf("%s/%s", httpLink, a.ApiKey)
	}
	if wsLink != "" {
		a.WebsocketLink = fmt.Sprintf("%s/%s", wsLink, a.ApiKey)
	}
	return a, a.save()
}

func (a *RpcApp) generateKey() error {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	str := fmt.Sprintf("%s-%s-%s", a.Account, a.Name, newUUID.String())
	a.ApiKey = fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return nil
}

func (a *RpcApp) save() error {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return err
	}
	// 获取 app_id
	var appID int64
	if err = db.Model(&RpcApp{}).Select("max(app_id)").Where("account = ?", a.Account).Scan(&appID).Error; err != nil {
		appID = 0
	}
	a.AppID = appID + 1
	return db.Model(&RpcApp{}).Create(a).Error
}

func deleteApp(account string, id int) error {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return err
	}
	return db.Delete(&RpcApp{}, "account = ? AND app_id = ?", account, id).Error
}

func getApp(account string, id int) (*RpcApp, error) {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return nil, err
	}
	var app RpcApp
	if err := db.Where("account = ? AND app_id = ?", account, id).First(&app).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

func getAppByName(account string, name string) (*ApiResponseRpcApp, error) {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return nil, err
	}
	var app RpcApp
	if err := db.Where("account = ? AND name = ?", account, name).First(&app).Error; err != nil {
		return nil, err
	}
	var appResp ApiResponseRpcApp
	appResp.RpcApp = &app
	appResp.TotalRequests24h, err = app.getTotalRequests24hWithStatusAll()
	if err != nil {
		logger.Errorf("getTotalRequestsTodayWithStatusAll err: %s", err)
	}
	appResp.DaylyRequests7Days, err = app.getDaylyRequests7DaysWithStatusAll()
	if err != nil {
		logger.Errorf("getDaylyRequests7DaysWithStatusAll err: %s", err)
	}
	appResp.CodeExamples = newCodeExample()
	return &appResp, nil
}

func accountHaveApp(account string, appKey string) bool {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return false
	}
	var app RpcApp
	if err := db.Where("account = ? AND api_key = ?", account, appKey).First(&app).Error; err != nil {
		return false
	}
	return true
}

func getAppBaseInfoByChainNetwork(account string, chain ChainType, network NetworkType) (*ApiResponseRpcApp, error) {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return nil, err
	}
	var app RpcApp
	if err := db.Where("account = ? AND chain = ? AND network = ?", account, chain.String(), network.String()).First(&app).Error; err != nil {
		return nil, err
	}
	var appResp ApiResponseRpcApp
	app.Account = ""
	appResp.RpcApp = &app
	appResp.CodeExamples = newCodeExample()
	return &appResp, nil
}

func getAppByChainNetwork(account string, chain ChainType, network NetworkType) (*ApiResponseRpcApp, error) {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return nil, err
	}
	var app RpcApp
	if err := db.Where("account = ? AND chain = ? AND network = ?", account, chain.String(), network.String()).First(&app).Error; err != nil {
		return nil, err
	}
	var appResp ApiResponseRpcApp
	appResp.RpcApp = &app
	appResp.TotalRequests24h, err = app.getTotalRequests24hWithStatusAll()
	if err != nil {
		logger.Errorf("getTotalRequestsTodayWithStatusAll err: %s", err)
	}
	appResp.DaylyRequests7Days, err = app.getDaylyRequests7DaysWithStatusAll()
	if err != nil {
		logger.Errorf("getDaylyRequests7DaysWithStatusAll err: %s", err)
	}
	appResp.TotalRequestsAll, err = app.getTotalRequestsAll()
	if err != nil {
		logger.Errorf("getTotalRequestsAll err: %s", err)
	}
	appResp.CodeExamples = newCodeExample()
	return &appResp, nil
}

func filterAppsMainnet(account string) ([]*ApiResponseRpcApp, error) {
	return filterApps(account, "mainnet")
}

func filterAppsTestnet(account string) ([]*ApiResponseRpcApp, error) {
	// 这里会返回所有的 testnet
	return filterApps(account, "testnet")
}

func filterApps(account string, network string) ([]*ApiResponseRpcApp, error) {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return nil, err
	}
	var apps []*RpcApp
	if network == "mainnet" {
		err = db.Model(&RpcApp{}).Where("account = ? AND network = 'Mainnet'", account).Order("id asc").Find(&apps).Error
	} else {
		err = db.Model(&RpcApp{}).Where("account = ? AND network != 'Mainnet'", account).Order("id asc").Find(&apps).Error
	}
	if err != nil {
		return nil, err
	}
	var apiResponseApps []*ApiResponseRpcApp
	for i := range apps {
		apps[i].Account = ""
		var appResp ApiResponseRpcApp
		appResp.RpcApp = apps[i]
		appResp.TotalRequests24h, err = apps[i].getTotalRequests24hWithStatusAll()
		if err != nil {
			logger.Errorf("getTotalRequestsTodayWithStatusAll err: %s", err)
		}
		appResp.DaylyRequests7Days, err = apps[i].getDaylyRequests7DaysWithStatusAll()
		if err != nil {
			logger.Errorf("getDaylyRequests7DaysWithStatusAll err: %s", err)
		}
		appResp.TotalRequestsAll, err = apps[i].getTotalRequestsAll()
		if err != nil {
			logger.Errorf("getTotalRequestsAll err: %s", err)
		}
		appResp.CodeExamples = newCodeExample()
		apiResponseApps = append(apiResponseApps, &appResp)
	}
	return apiResponseApps, nil
}

func getApps(account string) ([]*ApiResponseRpcApp, error) {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return nil, err
	}
	var apps []*RpcApp
	if err := db.Model(&RpcApp{}).Where("account = ?", account).Order("id asc").Find(&apps).Error; err != nil {
		return nil, err
	}
	var apiResponseApps []*ApiResponseRpcApp
	for i := range apps {
		apps[i].Account = ""
		var appResp ApiResponseRpcApp
		appResp.RpcApp = apps[i]
		appResp.TotalRequests24h, err = apps[i].getTotalRequests24hWithStatusAll()
		if err != nil {
			logger.Errorf("getTotalRequestsTodayWithStatusAll err: %s", err)
		}
		appResp.DaylyRequests7Days, err = apps[i].getDaylyRequests7DaysWithStatusAll()
		if err != nil {
			logger.Errorf("getDaylyRequests7DaysWithStatusAll err: %s", err)
		}
		appResp.TotalRequestsAll, err = apps[i].getTotalRequestsAll()
		if err != nil {
			logger.Errorf("getTotalRequestsAll err: %s", err)
		}
		appResp.CodeExamples = newCodeExample()
		apiResponseApps = append(apiResponseApps, &appResp)
	}
	return apiResponseApps, nil
}

func getAppsPagination(account string, p *Pagination) ([]*ApiResponseRpcApp, *Pagination, error) {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return nil, p, err
	}
	var apps []*RpcApp
	if err := db.Model(&RpcApp{}).Limit(p.Size).Offset(p.Size*(p.Page-1)).Where("account = ?", account).Order("id asc").Find(&apps).Error; err != nil {
		return nil, p, err
	}
	var apiResponseApps []*ApiResponseRpcApp
	for i := range apps {
		apps[i].Account = ""
		var appResp ApiResponseRpcApp
		appResp.RpcApp = apps[i]
		appResp.TotalRequests24h, err = apps[i].getTotalRequests24hWithStatusAll()
		if err != nil {
			logger.Errorf("getTotalRequestsTodayWithStatusAll err: %s", err)
		}
		appResp.DaylyRequests7Days, err = apps[i].getDaylyRequests7DaysWithStatusAll()
		if err != nil {
			logger.Errorf("getDaylyRequests7DaysWithStatusAll err: %s", err)
		}
		appResp.TotalRequestsAll, err = apps[i].getTotalRequestsAll()
		if err != nil {
			logger.Errorf("getTotalRequestsAll err: %s", err)
		}
		appResp.CodeExamples = newCodeExample()
		apiResponseApps = append(apiResponseApps, &appResp)
	}
	var count int64
	if err := db.Model(&RpcApp{}).Where("account = ?", account).Count(&count).Error; err != nil {
		return nil, p, err
	}
	p.Total = count
	return apiResponseApps, p, nil
}

// 获取总请求数
func (a *RpcApp) getTotalRequests(statusFilter string) (int64, error) {
	meili, err := application.GetBean[*meilisearch.Client]("meiliSearch")
	if err != nil {
		return 0, err
	}
	resp, err := meili.Index("nginx").Search(a.ApiKey, &meilisearch.SearchRequest{
		Filter: statusFilter,
	})
	if err != nil {
		logger.Errorf("git total requests all error: %s", err)
		return 0, err
	}
	return resp.EstimatedTotalHits, nil
}

// 获取总请求数状态为 200
func (a *RpcApp) getTotalRequests200() (int64, error) {
	return a.getTotalRequests("status = 200")
}

// 获取总请求数状态为全部
func (a *RpcApp) getTotalRequestsAll() (int64, error) {
	return a.getTotalRequests("")
}

// 获取总请求数状态不为 200
func (a *RpcApp) getTotalRequestsNot200() (int64, error) {
	return a.getTotalRequests("status != 200")
}

// 获取今日请求数
func (a *RpcApp) getTotalRequests24h(statusFilter string) (int64, error) {
	meili, err := application.GetBean[*meilisearch.Client]("meiliSearch")
	if err != nil {
		logger.Errorf("getTotalRequestsToday err: %s", err)
		return 0, err
	}
	twentyFourHoursAgo := float64(time.Now().Add(-24 * time.Hour).Unix())
	resp, err := meili.Index("nginx").Search(a.ApiKey, &meilisearch.SearchRequest{
		Filter: []string{statusFilter, fmt.Sprintf("msec >= %f", twentyFourHoursAgo)},
	})
	if err != nil {
		logger.Errorf("getTotalRequestsToday err: %s", err)
		return 0, err
	}
	return resp.EstimatedTotalHits, nil
}

func (a *RpcApp) getTotalRequests24hWithStatus200() (int64, error) {
	return a.getTotalRequests24h("status = 200")
}

func (a *RpcApp) getTotalRequests24hWithStatusAll() (int64, error) {
	return a.getTotalRequests24h("")
}

func (a *RpcApp) getTotalRequests24hWithStatusIsNot200() (int64, error) {
	return a.getTotalRequests24h("status != 200")
}

// 获取 7 天请求数
func (a *RpcApp) getDaylyRequests7Days(statusFilter string) ([]DateRequest, error) {
	meili, err := application.GetBean[*meilisearch.Client]("meiliSearch")
	if err != nil {
		return nil, err
	}
	// 使用美丽搜索过滤过去 7 天的数据，并且状态码为 200 的数据
	// 以自然日为单位，即 2021-01-01 00:00:00 到 2021-01-01 23:59:59
	// 不包括今天
	// 先计算昨天的结束时间，然后计算前 6 天的开始时间和结束时间
	nowUnix := time.Now().Unix()
	yesterdayStartUnix := nowUnix - int64(nowUnix%86400) - 86400
	yesterdayEndUnix := yesterdayStartUnix + 86400 - 1
	var result []DateRequest
	for i := 0; i < 7; i++ {
		startDate := yesterdayStartUnix - int64((i)*86400)
		endDate := yesterdayEndUnix - int64((i)*86400)
		resp, err := meili.Index("nginx").Search(a.ApiKey, &meilisearch.SearchRequest{
			Filter: []string{fmt.Sprintf("msec >= %d", startDate), fmt.Sprintf("msec < %d", endDate), statusFilter},
		})
		if err != nil {
			return nil, err
		}
		dateRequest := DateRequest{
			StartTime:     time.Unix(startDate, 0).UTC().Format(time.RFC3339),
			EndTime:       time.Unix(endDate, 0).UTC().Format(time.RFC3339),
			RequestNumber: resp.EstimatedTotalHits,
		}
		result = append(result, dateRequest)
	}
	return reverse(result), nil
}

func reverse(in []DateRequest) []DateRequest {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
	return in
}

// 获取 7 天请求数状态码为 200 的数据
func (a *RpcApp) getDaylyRequests7DaysWithStatus200() ([]DateRequest, error) {
	return a.getDaylyRequests7Days("status = 200")
}

// 获取 7 天请求数所有状态
func (a *RpcApp) getDaylyRequests7DaysWithStatusAll() ([]DateRequest, error) {
	return a.getDaylyRequests7Days("")
}
func (a *RpcApp) getDaylyRequests7DaysWithStatusIsNot200() ([]DateRequest, error) {
	return a.getDaylyRequests7Days("status != 200")
}

type RpcAppRequestLog struct {
	Number        int64  `json:"number"`
	Time          string `json:"time"`
	RequestEvent  string `json:"request_event"`
	RequestResult string `json:"request_result"`
}

type NginxAccessLog struct {
	Msec                   float64 `json:"msec"`
	Connection             string  `json:"connection"`
	ConnectionRequests     string  `json:"connection_requests"`
	Pid                    string  `json:"pid"`
	RequestID              string  `json:"request_id"`
	RequestLength          string  `json:"request_length"`
	RemoteAddr             string  `json:"remote_addr"`
	RemoteUser             string  `json:"remote_user"`
	RemotePort             string  `json:"remote_port"`
	TimeLocal              string  `json:"time_local"`
	TimeISO8601            string  `json:"time_iso8601"`
	Request                string  `json:"request"`
	RequestURI             string  `json:"request_uri"`
	Args                   string  `json:"args"`
	Status                 string  `json:"status"`
	BodyBytesSent          string  `json:"body_bytes_sent"`
	BytesSent              string  `json:"bytes_sent"`
	HttpReferer            string  `json:"http_referer"`
	HttpUserAgent          string  `json:"http_user_agent"`
	HttpXForwardedFor      string  `json:"http_x_forwarded_for"`
	HttpHost               string  `json:"http_host"`
	ServerName             string  `json:"server_name"`
	RequestTime            string  `json:"request_time"`
	Upstream               string  `json:"upstream"`
	UpstreamConnectTime    string  `json:"upstream_connect_time"`
	UpstreamHeaderTime     string  `json:"upstream_header_time"`
	UpstreamResponseTime   string  `json:"upstream_response_time"`
	UpstreamResponseLength string  `json:"upstream_response_length"`
	UpstreamCacheStatus    string  `json:"upstream_cache_status"`
	SslProtocol            string  `json:"ssl_protocol"`
	SslCipher              string  `json:"ssl_cipher"`
	Scheme                 string  `json:"scheme"`
	RequestMethod          string  `json:"request_method"`
	ServerProtocol         string  `json:"server_protocol"`
	Pipe                   string  `json:"pipe"`
	GzipRatio              string  `json:"gzip_ratio"`
	HttpCFRay              string  `json:"http_cf_ray"`
}

func (a *RpcApp) getAppRequestLogs(appKey string, p Pagination) ([]*RpcAppRequestLog, *Pagination, error) {
	meili, err := application.GetBean[*meilisearch.Client]("meiliSearch")
	if err != nil {
		return nil, &p, err
	}
	resp, err := meili.Index("nginx").Search(a.ApiKey, &meilisearch.SearchRequest{
		Limit:  int64(p.Size),
		Offset: int64(p.Size * (p.Page - 1)),
		Sort:   []string{"msec:desc"},
	})
	if err != nil {
		return nil, &p, err
	}
	p.Total = resp.EstimatedTotalHits
	var logs []*RpcAppRequestLog
	for i := range resp.Hits {
		log := &RpcAppRequestLog{}
		log.Number = resp.EstimatedTotalHits - int64(i) - int64(p.Size*(p.Page-1))
		log.Time = resp.Hits[i].(map[string]any)["time_iso8601"].(string)
		log.RequestEvent = resp.Hits[i].(map[string]any)["args"].(string)
		log.RequestResult = resp.Hits[i].(map[string]any)["status"].(string)
		logs = append(logs, log)
	}
	return logs, &p, nil
}
