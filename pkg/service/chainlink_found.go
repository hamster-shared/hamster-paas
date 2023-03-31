package service

import "gorm.io/gorm"

type ChainLinkFoundService struct {
	db *gorm.DB
}

func NewChainLinkFoundService(db *gorm.DB) *ChainLinkFoundService {
	return &ChainLinkFoundService{
		db: db,
	}
}

func (c *ChainLinkFoundService) AddFundsForSubscription(SubscriptionId int64, incr float64) {
}
