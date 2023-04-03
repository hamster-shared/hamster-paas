package models

import (
	"fmt"
	"hamster-paas/pkg/application"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	db         *gorm.DB
	Address    string `json:"address"`
	AppIdIndex int
}

func NewAccount(address string) (*Account, error) {
	a := &Account{
		Address:    address,
		AppIdIndex: 0,
	}
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %s", err)
	}
	a.db = db
	return a, a.save()
}

func (a *Account) save() error {
	return a.db.Create(a).Error
}

func GetAccount(address string) (*Account, error) {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %s", err)
	}
	var account Account
	if err := db.Where("address = ?", address).First(&account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return NewAccount(address)
		}
		return nil, fmt.Errorf("failed to get account: %s", err)
	}
	account.db = db
	return &account, nil
}

func (a *Account) CreateApp(name, description string, chain ChainType, network NetworkType) (*App, error) {
	return NewApp(a.Address, a.AppIdIndex+1, name, description, chain, network)
}

func (a *Account) DeleteApp(id int) error {
	return DeleteApp(a.Address, id)
}

func (a *Account) GetApps(p Pagination) ([]*ApiResponseApp, Pagination, error) {
	return GetApps(a.Address, p)
}

func (a *Account) GetApp(id int) (*App, error) {
	return GetApp(a.Address, id)
}
