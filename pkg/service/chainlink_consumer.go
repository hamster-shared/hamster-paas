package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/models/vo"
)

type ChainLinkConsumerService struct {
	db *gorm.DB
}

func NewChainLinkConsumerService(db *gorm.DB) *ChainLinkConsumerService {
	return &ChainLinkConsumerService{
		db: db,
	}
}

func (c *ChainLinkConsumerService) CreateConsumer(consumer models.Consumer, subscriptionId int) (int64, error) {
	var isExited int64
	err := c.db.Model(models.Consumer{}).Where("subscription_id = ? AND consumer_address = ?", subscriptionId, consumer.ConsumerAddress).Count(&isExited).Error
	if err == gorm.ErrRecordNotFound {
		c.db.Create(&consumer)
		var count int64
		c.db.Model(models.Consumer{}).Where("subscription_id = ?", subscriptionId).Count(&count)
		return count, nil
	}
	return 0, errors.New(fmt.Sprintf("consumer address :%s already exists in subscription id: %d", consumer.ConsumerAddress, subscriptionId))
}

// ConsumerList get consumer list
func (c *ChainLinkConsumerService) ConsumerList(subscriptionId, page, size int, userId int64) (*vo.ChainLinkConsumerPage, error) {
	var total int64
	var chainLinkConsumerPage vo.ChainLinkConsumerPage
	var chainLinkConsumerList []models.Consumer
	var chainLinkConsumerVoList []vo.ChainLinkConsumerVo
	tx := c.db.Model(models.Consumer{}).Where("user_id = ? and subscription_id", userId, subscriptionId)
	result := tx.Order("created DESC").Offset((page - 1) * size).Limit(size).Find(&chainLinkConsumerList).Offset(-1).Limit(-1).Count(&total)
	if result.Error != nil {
		return &chainLinkConsumerPage, result.Error
	}
	copier.Copy(&chainLinkConsumerVoList, &chainLinkConsumerList)
	chainLinkConsumerPage.Data = chainLinkConsumerVoList
	chainLinkConsumerPage.Total = total
	chainLinkConsumerPage.Page = page
	chainLinkConsumerPage.PageSize = size
	return &chainLinkConsumerPage, nil
}

// DeleteConsumer delete consumer by id
func (c *ChainLinkConsumerService) DeleteConsumer(subscriptionId, consumerId int64) error {
	err := c.db.Transaction(func(tx *gorm.DB) error {
		if err := c.db.Debug().Where("id = ? and subscription_id = ", consumerId, subscriptionId).Delete(&models.Consumer{}).Error; err != nil {
			return err
		}
		var subscriptionData models.Subscription
		if err := c.db.Where("id = ?", subscriptionId).First(&subscriptionData).Error; err != nil {
			return err
		}
		if subscriptionData.Consumers > 0 {
			subscriptionData.Consumers = subscriptionData.Consumers - 1
			if err := c.db.Save(&subscriptionData).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
