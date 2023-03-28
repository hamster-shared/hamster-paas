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
	appIdIndex int
}

func NewAccount(address string) (*Account, error) {
	a := &Account{
		Address:    address,
		appIdIndex: 0,
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
