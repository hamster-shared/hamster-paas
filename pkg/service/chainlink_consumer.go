package service

import (
	"gorm.io/gorm"
	"hamster-paas/pkg/models"
)

type ChainLinkConsumerService struct {
	db *gorm.DB
}

func NewChainLinkConsumerService(db *gorm.DB) *ChainLinkConsumerService {
	return &ChainLinkConsumerService{
		db: db,
	}
}

func (c *ChainLinkConsumerService) CreateConsumer(consumer models.Consumer, subscriptionId int) int64 {
	c.db.Create(&consumer)
	var count int64
	c.db.Table("t_cl_consumer").Where("subscription_id = ?", subscriptionId).Count(&count)
	return count
}
