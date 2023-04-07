package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"hamster-paas/pkg/consts"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/models/vo"
	"hamster-paas/pkg/utils/logger"
	"time"
)

type ChainLinkDepositService struct {
	db *gorm.DB
}

func NewChainLinkDepositService(db *gorm.DB) *ChainLinkDepositService {
	return &ChainLinkDepositService{
		db: db,
	}
}

// DepositList  query chain link deposit list by subscription id
func (d *ChainLinkDepositService) DepositList(subscriptionId, page, size int, userId int64) (*vo.ChainLinkDepositPage, error) {
	var total int64
	var chainLinkDepositPage vo.ChainLinkDepositPage
	var chainLinkDepositList []models.Deposit
	var chainLinkDepositVoList []vo.ChainLinkDepositVo
	tx := d.db.Model(models.Deposit{}).Where("user_id = ? and subscription_id = ? ", userId, subscriptionId)
	result := tx.Order("created DESC").Offset((page - 1) * size).Limit(size).Find(&chainLinkDepositList).Offset(-1).Limit(-1).Count(&total)
	if result.Error != nil {
		return &chainLinkDepositPage, result.Error
	}
	copier.Copy(&chainLinkDepositVoList, &chainLinkDepositList)
	chainLinkDepositPage.Data = chainLinkDepositVoList
	chainLinkDepositPage.Total = total
	chainLinkDepositPage.Page = page
	chainLinkDepositPage.PageSize = size
	return &chainLinkDepositPage, nil
}

// AddDeposit TODO 需要异步检查
func (d *ChainLinkDepositService) AddDeposit(subscriptionId int64, consumerAddress string, incr float64, transactionTx string, userId int64, subscriptionService ChainLinkSubscriptionService, poolService PoolService) error {
	// 检查该id是否存在且success
	_, err := subscriptionService.GetSubscriptionById(int(subscriptionId))
	if err != nil {
		return err
	}
	var deposit models.Deposit
	deposit.SubscriptionId = subscriptionId
	deposit.Created = time.Now()
	deposit.ConsumerAddress = consumerAddress
	deposit.Amount = incr
	deposit.TransactionTx = transactionTx
	deposit.UserId = uint64(userId)
	deposit.Status = consts.PENDING
	d.db.Model(models.Deposit{}).Create(&deposit)
	// TODO 异步检查
	poolService.Submit(func() {
		time.Sleep(time.Second * 20)
		d.db.Model(models.Deposit{}).Where("id = ?", deposit.Id).Update("status", consts.SUCCESS)
		logger.Infof("add deposit: %d Tx valid, change deposit status to success", deposit.Id)
	})
	return nil
}

func (d *ChainLinkDepositService) UpdateDepositStatus(userId uint64, param vo.ChainLinkFoundUpdateParam) error {
	//获取id对应的记录
	var deposit models.Deposit
	d.db.Model(models.Deposit{}).Where("id = ?", param.Id).First(&deposit)
	// 判断该deposit是否是符合要求

	if deposit.TransactionTx == param.TransactionTx && deposit.ConsumerAddress == param.ConsumerAddress && deposit.UserId == userId && param.SubscriptionId == deposit.SubscriptionId {
		d.db.Model(models.Deposit{}).Where("id = ?", param.Id).Update("status", param.NewStatus)
		return nil
	}
	return errors.New(fmt.Sprintf("consumer id :%s not valid, other col not confirm", param.Id))
}
