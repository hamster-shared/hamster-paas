package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"hamster-paas/pkg/consts"
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

// CreateConsumer
// @param consumer中的subscription id指的是subscription表主键id
// TODO: 需要监听链更改状态
func (c *ChainLinkConsumerService) CreateConsumer(consumer models.Consumer, subscriptionService ChainLinkSubscriptionService, poolService PoolService) (int64, error) {
	// 确认subscription存在
	_, err := subscriptionService.GetSubscriptionById(int(consumer.SubscriptionId))
	if err != nil {
		return -1, err
	}
	var isExited int64
	// 判断该consumer是否存在
	c.db.Model(models.Consumer{}).Where("subscription_id = ? AND consumer_address = ? AND (status = ? OR status = ?)", consumer.SubscriptionId, consumer.ConsumerAddress, consts.SUCCESS, consts.PENDING).Count(&isExited)
	if isExited > 0 {
		return -1, errors.New(fmt.Sprintf("consumer address :%s already exists in subscription id: %d", consumer.ConsumerAddress, consumer.SubscriptionId))
	}
	// 不存在就创建
	c.db.Create(&consumer)
	return consumer.Id, nil
}

// ConsumerList get consumer list
func (c *ChainLinkConsumerService) ConsumerList(subscriptionId, page, size int, userId int64) (*vo.ChainLinkConsumerPage, error) {
	var total int64
	var chainLinkConsumerPage vo.ChainLinkConsumerPage
	var chainLinkConsumerList []models.Consumer
	var chainLinkConsumerVoList []vo.ChainLinkConsumerVo
	tx := c.db.Model(models.Consumer{}).Where("user_id = ? and subscription_id = ?", userId, subscriptionId)
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
		if err := c.db.Debug().Where("id = ? and subscription_id = ?", consumerId, subscriptionId).Delete(&models.Consumer{}).Error; err != nil {
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

func (c *ChainLinkConsumerService) ChangeConsumerStatus(param vo.ChainLinkConsumerUpdateParam, userId uint64) error {
	//获取id对应的记录
	var consumer models.Consumer
	c.db.Model(models.Consumer{}).Where("id = ?", param.Id).First(&consumer)
	// 如果已经是成功状态，不做操作
	if consumer.Status == consts.SUCCESS {
		return nil
	}
	// 判断该consumer是否是符合要求
	if consumer.TransactionTx == param.TransactionTx && consumer.ConsumerAddress == param.ConsumerAddress && consumer.UserId == userId && param.SubscriptionId == consumer.SubscriptionId {
		c.db.Model(models.Consumer{}).Where("id = ?", param.Id).Update("status", param.NewStatus)
		return nil
	}
	return errors.New(fmt.Sprintf("consumer id :%s not valid, other col not confirm", param.Id))
}
