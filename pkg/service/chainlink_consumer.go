package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"hamster-paas/pkg/consts"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/models/vo"
	"hamster-paas/pkg/rpc/eth"
	"hamster-paas/pkg/utils/logger"
	"time"
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
	subscription, err := subscriptionService.GetSubscriptionById(int(consumer.SubscriptionId))
	if err != nil {
		return -1, err
	}
	network, err := models.ParseNetworkType(subscription.Network)
	if err != nil {
		logger.Error(fmt.Sprintf("network format error: %s", err.Error()))
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
	// 异步判断TX status
	poolService.Submit(func() {
		checkAndChangeConsumerStatus(network, consumer, c.db, subscriptionService)
	})
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

func (c *ChainLinkConsumerService) ConsumerAddressList(subscriptionId, userId int64) ([]string, error) {
	var data []string
	res := c.db.Model(models.Consumer{}).Distinct("consumer_address").Select("consumer_address").Where("subscription_id=? and user_id=?", subscriptionId, userId).Find(&data)
	if res.Error != nil {
		return data, res.Error
	}
	return data, nil
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

func (c *ChainLinkConsumerService) ChangeConsumerStatus(param vo.ChainLinkConsumerUpdateParam, userId uint64, subscriptionService ChainLinkSubscriptionService) error {
	//获取id对应的记录
	var consumer models.Consumer
	c.db.Model(models.Consumer{}).Where("id = ?", param.Id).First(&consumer)
	// 如果已经是成功状态，不做操作
	if consumer.Status == param.NewStatus {
		return nil
	}
	// 判断该consumer是否是符合要求,符合要求，修改status，增加subscription的consumer数量
	if consumer.TransactionTx == param.TransactionTx && consumer.ConsumerAddress == param.ConsumerAddress && consumer.UserId == userId && param.SubscriptionId == consumer.SubscriptionId {
		c.db.Model(models.Consumer{}).Where("id = ?", param.Id).Update("status", param.NewStatus)
		// 获取指定id的subscription
		subscription, err := subscriptionService.GetSubscriptionById(int(param.SubscriptionId))
		if err != nil {
			return fmt.Errorf("subscription id not exit")
		}
		// 更新subscription的consumer数量
		err = subscriptionService.UpdateConsumerNums(uint(param.SubscriptionId), int64(subscription.Consumers+1))
		if err != nil {
			return fmt.Errorf("update consumer nums failed: %s", err.Error())
		}
		return nil
	}
	return errors.New(fmt.Sprintf("consumer id :%d not valid, other col not confirm", param.Id))
}

// 用于检查tx的状态，并且修改deposit的status
func checkAndChangeConsumerStatus(network models.NetworkType, consumer models.Consumer, db *gorm.DB, subscriptionService ChainLinkSubscriptionService) {
	client := eth.GetChainClient(network.NetworkType())
	if client == nil {
		return
	}
	times := 0
	needFalid := false
	for {
		if times == 6 {
			needFalid = true
			break
		}
		time.Sleep(time.Second * 20)
		times++
		// 拿到数据库中状态,判断是否要主动结束轮询
		var c_ models.Consumer
		db.Model(models.Consumer{}).Where("id = ?", consumer.Id).First(&c_)
		// status == Success， 主动结束轮询
		if c_.Status == consts.SUCCESS {
			break
		}
		// 获取tx状态
		txStatus, err := eth.GetTxStatus(consumer.TransactionTx, network.NetworkType(), client)
		if err != nil {
			logger.Errorf("get tx status error: %s", err)
			continue
		}
		if txStatus == 1 {
			// 修改状态为成功
			logger.Infof("Create Consumer : Tx Success, change consumer id: %d status to success", consumer.Id)
			db.Model(models.Consumer{}).Where("id = ?", consumer.Id).Update("status", consts.SUCCESS)
			// 获取指定id的subscription
			subscription, err := subscriptionService.GetSubscriptionById(int(consumer.SubscriptionId))
			if err != nil {
				logger.Infof("check and change consumer status error, subscription id: %d not exit", consumer.SubscriptionId)
				break
			}
			// 更新subscription的consumer数量
			err = subscriptionService.UpdateConsumerNums(uint(consumer.SubscriptionId), int64(subscription.Consumers+1))
			if err != nil {
				logger.Infof("check and change consumer status error, update subscription id: %d consumers error", consumer.SubscriptionId)
				break
			}
			break
		} else if txStatus == 0 {
			// 修改状态为失败
			logger.Infof("Create Consumer : Tx failed, change consumer id: %d status to failed", consumer.Id)
			db.Model(models.Consumer{}).Where("id = ?", consumer.Id).Update("status", consts.FAILED)
			break
		}
	}
	if needFalid {
		// 更新状态为失败
		logger.Infof("Create consumer : Query timeout, change consumer id: %d status to failed", consumer.Id)
		db.Model(models.Consumer{}).Where("id = ?", consumer.Id).Update("status", consts.FAILED)
	}
}
