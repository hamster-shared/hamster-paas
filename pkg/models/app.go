package models

import (
	"hamster-paas/pkg/application"

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

func NewApp(account string, id int, name, description string, chain ChainType, network NetworkType) (*App, error) {
	a := &App{
		Account:     account,
		AppId:       id,
		Name:        name,
		Description: description,
		Chain:       chain.String(),
		Network:     network.String(),
	}
	return a, a.save()
}

func (a *App) save() error {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return err
	}
	return db.Create(a).Error
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

func GetApps(account string, pagination ApiRequestPagination) ([]*App, Pagination, error) {
	var p Pagination
	p.Page = pagination.Page
	p.Size = pagination.Size
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return nil, p, err
	}
	var apps []*App
	if err := db.Where("account = ?", account).Order("app_id desc").Limit(pagination.Size).Offset((pagination.Page - 1) * pagination.Size).Find(&apps).Error; err != nil {
		return nil, p, err
	}
	db.Model(&App{}).Where("account = ?", account).Count(&p.Total)
	return apps, p, nil
}
