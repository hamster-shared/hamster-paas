package service

import (
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

func (c *ChainLinkConsumerService) CreateConsumer(consumer models.Consumer, subscriptionId int) int64 {
	c.db.Create(&consumer)
	var count int64
	c.db.Table("t_cl_consumer").Where("subscription_id = ?", subscriptionId).Count(&count)
	return count
}

// ConsumerList get consumer list
func (c *ChainLinkConsumerService) ConsumerList(page, size int, userId int64) (*vo.ChainLinkConsumerPage, error) {
	var total int64
	var chainLinkConsumerPage vo.ChainLinkConsumerPage
	var chainLinkConsumerList []models.Consumer
	var chainLinkConsumerVoList []vo.ChainLinkConsumerVo
	tx := c.db.Model(models.Consumer{}).Where("user_id = ?", userId)
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
func (c *ChainLinkConsumerService) DeleteConsumer(id int64) error {
	err := c.db.Debug().Where("id = ? ", id).Delete(&models.Consumer{}).Error
	if err != nil {
		return err
	}
	return nil
}
