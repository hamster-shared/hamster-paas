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

type App struct {
	AppId         int    `json:"app_id"`
	Account       string `json:"account"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Chain         string `json:"chain"`
	Network       string `json:"network"`
	ApiKey        string `json:"api_key"`
	HttpLink      string `json:"http_link"`
	WebsocketLink string `json:"websocket_link"`
}

type ApiResponseApp struct {
	*App
	CodeExamples       []CodeExample `json:"code_examples"`
	TotalRequestsToday int64         `json:"total_requests_today"`
	DaylyRequests7Days []int64       `json:"dayly_requests_7days"`
}

type CodeExample struct{}

func NewApp(account string, id int, name, description string, chain ChainType, network NetworkType) (*App, error) {
	a := &App{
		Account:     account,
		AppId:       id,
		Name:        name,
		Description: description,
		Chain:       chain.String(),
		Network:     network.String(),
	}
	err := a.generateKey()
	if err != nil {
		return nil, err
	}
	return a, a.save()
}

func (a *App) generateKey() error {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	str := fmt.Sprintf("%s-%d-%s", a.Account, a.AppId, newUUID.String())
	a.ApiKey = fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return nil
}

func (a *App) save() error {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return err
	}
	if err := db.Model(&Account{}).Where("address = ?", a.Account).Update("app_id_index", gorm.Expr("app_id_index + ?", 1)).Error; err != nil {
		return err
	}
	return db.Model(&App{}).Create(a).Error
}

func DeleteApp(account string, id int) error {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return err
	}
	return db.Delete(&App{}, "account = ? AND app_id = ?", account, id).Error
}

func GetApp(account string, id int) (*App, error) {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return nil, err
	}
	var app App
	if err := db.Where("account = ? AND app_id = ?", account, id).First(&app).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

func GetApps(account string, pagination Pagination) ([]*ApiResponseApp, Pagination, error) {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return nil, pagination, err
	}
	var apps []*App
	limit := pagination.Size
	offset := (pagination.Page - 1) * pagination.Size
	if err := db.Model(&App{}).Where("account = ?", account).Order("app_id desc").Limit(limit).Offset(offset).Find(&apps).Error; err != nil {
		return nil, pagination, err
	}
	db.Model(&App{}).Where("account = ?", account).Count(&pagination.Total)
	var apiResponseApps []*ApiResponseApp
	for i := range apps {
		var appResp ApiResponseApp
		appResp.App = apps[i]
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
	return apiResponseApps, pagination, nil
}

// 获取总请求数
func (a *App) getTotalRequests(statusFilter string) (int64, error) {
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
func (a *App) getTotalRequests200() (int64, error) {
	return a.getTotalRequests("status = 200")
}

// 获取总请求数状态为全部
func (a *App) getTotalRequestsAll() (int64, error) {
	return a.getTotalRequests("")
}

// 获取总请求数状态不为 200
func (a *App) getTotalRequestsNot200() (int64, error) {
	return a.getTotalRequests("status != 200")
}

// 获取今日请求数
func (a *App) getTotalRequestsToday(statusFilter string) (int64, error) {
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

func (a *App) getTotalRequestsTodayWithStatus200() (int64, error) {
	return a.getTotalRequestsToday("status = 200")
}

func (a *App) getTotalRequestsTodayWithStatusAll() (int64, error) {
	return a.getTotalRequestsToday("")
}

func (a *App) getTotalRequestsTodayWithStatusIsNot200() (int64, error) {
	return a.getTotalRequestsToday("status != 200")
}

// 获取 7 天请求数
func (a *App) getDaylyRequests7Days(statusFilter string) ([]int64, error) {
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
	return result, nil
}

// 获取 7 天请求数状态码为 200 的数据
func (a *App) getDaylyRequests7DaysWithStatus200() ([]int64, error) {
	return a.getDaylyRequests7Days("status = 200")
}

// 获取 7 天请求数所有状态
func (a *App) getDaylyRequests7DaysWithStatusAll() ([]int64, error) {
	return a.getDaylyRequests7Days("")
}
func (a *App) getDaylyRequests7DaysWithStatusIsNot200() ([]int64, error) {
	return a.getDaylyRequests7Days("status != 200")
}
