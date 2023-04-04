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
	Account       string `json:"account"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Chain         string `json:"chain"`
	Network       string `json:"network"`
	ApiKey        string `json:"api_key"`
	HttpLink      string `json:"http_link"`
	WebsocketLink string `json:"websocket_link"`
}

type ApiResponseRpcApp struct {
	*RpcApp
	CodeExamples       []RpcCodeExample `json:"code_examples"`
	TotalRequestsToday int64            `json:"total_requests_today"`
	DaylyRequests7Days []int64          `json:"dayly_requests_7days"`
}

type RpcCodeExample struct{}

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
	appResp.TotalRequestsToday, err = app.getTotalRequestsTodayWithStatusAll()
	if err != nil {
		logger.Errorf("getTotalRequestsTodayWithStatusAll err: %s", err)
	}
	appResp.DaylyRequests7Days, err = app.getDaylyRequests7DaysWithStatusAll()
	if err != nil {
		logger.Errorf("getDaylyRequests7DaysWithStatusAll err: %s", err)
	}
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
	appResp.TotalRequestsToday, err = app.getTotalRequestsTodayWithStatusAll()
	if err != nil {
		logger.Errorf("getTotalRequestsTodayWithStatusAll err: %s", err)
	}
	appResp.DaylyRequests7Days, err = app.getDaylyRequests7DaysWithStatusAll()
	if err != nil {
		logger.Errorf("getDaylyRequests7DaysWithStatusAll err: %s", err)
	}
	return &appResp, nil
}

func getApps(account string) ([]*ApiResponseRpcApp, error) {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return nil, err
	}
	var apps []*RpcApp
	if err := db.Model(&RpcApp{}).Where("account = ?", account).Order("id desc").Find(&apps).Error; err != nil {
		return nil, err
	}
	var apiResponseApps []*ApiResponseRpcApp
	for i := range apps {
		var appResp ApiResponseRpcApp
		appResp.RpcApp = apps[i]
		appResp.TotalRequestsToday, err = apps[i].getTotalRequestsTodayWithStatusAll()
		if err != nil {
			logger.Errorf("getTotalRequestsTodayWithStatusAll err: %s", err)
		}
		appResp.DaylyRequests7Days, err = apps[i].getDaylyRequests7DaysWithStatusAll()
		if err != nil {
			logger.Errorf("getDaylyRequests7DaysWithStatusAll err: %s", err)
		}
		apiResponseApps = append(apiResponseApps, &appResp)
	}
	return apiResponseApps, nil
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
func (a *RpcApp) getTotalRequestsToday(statusFilter string) (int64, error) {
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
	logger.Debugf("getTotalRequestsToday resp: status: %s,time: %.3f, count: %d", statusFilter, twentyFourHoursAgo, resp.EstimatedTotalHits)
	return resp.EstimatedTotalHits, nil
}

func (a *RpcApp) getTotalRequestsTodayWithStatus200() (int64, error) {
	return a.getTotalRequestsToday("status = 200")
}

func (a *RpcApp) getTotalRequestsTodayWithStatusAll() (int64, error) {
	return a.getTotalRequestsToday("")
}

func (a *RpcApp) getTotalRequestsTodayWithStatusIsNot200() (int64, error) {
	return a.getTotalRequestsToday("status != 200")
}

// 获取 7 天请求数
func (a *RpcApp) getDaylyRequests7Days(statusFilter string) ([]int64, error) {
	meili, err := application.GetBean[*meilisearch.Client]("meiliSearch")
	if err != nil {
		return nil, err
	}
	// 使用美丽搜索过滤过去 7 天的数据，并且状态码为 200 的数据
	nowUnix := time.Now().Unix()
	var result []int64
	for i := 0; i < 7; i++ {
		startDate := nowUnix - int64((i+1)*86400)
		endDate := nowUnix - int64(i*86400)
		resp, err := meili.Index("nginx").Search(a.ApiKey, &meilisearch.SearchRequest{
			Filter: []string{fmt.Sprintf("msec >= %d", startDate), fmt.Sprintf("msec < %d", endDate), statusFilter},
		})
		if err != nil {
			return nil, err
		}
		result = append(result, resp.EstimatedTotalHits)
	}
	return reverse(result), nil
}

func reverse(in []int64) []int64 {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
	return in
}

// 获取 7 天请求数状态码为 200 的数据
func (a *RpcApp) getDaylyRequests7DaysWithStatus200() ([]int64, error) {
	return a.getDaylyRequests7Days("status = 200")
}

// 获取 7 天请求数所有状态
func (a *RpcApp) getDaylyRequests7DaysWithStatusAll() ([]int64, error) {
	return a.getDaylyRequests7Days("")
}
func (a *RpcApp) getDaylyRequests7DaysWithStatusIsNot200() ([]int64, error) {
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
