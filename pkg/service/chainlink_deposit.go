package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"hamster-paas/pkg/consts"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/models/vo"
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
func (d *ChainLinkDepositService) AddDeposit(subscriptionId int64, incr float64, transactionTx string, userId int64, subscriptionService ChainLinkSubscriptionService, poolService PoolService) (int64, error) {
	// 检查该id是否存在且success
	_, err := subscriptionService.GetSubscriptionById(int(subscriptionId))
	if err != nil {
		return -1, err
	}
	var deposit models.Deposit
	deposit.SubscriptionId = subscriptionId
	deposit.Created = time.Now()
	deposit.Amount = incr
	deposit.TransactionTx = transactionTx
	deposit.UserId = uint64(userId)
	deposit.Status = consts.PENDING
	err = d.db.Model(models.Deposit{}).Create(&deposit).Error
	if err != nil {
		return -1, err
	}
	return deposit.Id, nil
}

func (d *ChainLinkDepositService) UpdateDepositStatus(userId uint64, param vo.ChainLinkFoundUpdateParam) error {
	//获取id对应的记录
	var deposit models.Deposit
	d.db.Model(models.Deposit{}).Where("id = ?", param.Id).First(&deposit)
	// 如果deposit已经是成功状态，不做操作
	if deposit.Status == consts.SUCCESS {
		return nil
	}
	// 判断该deposit是否是符合要求
	if deposit.TransactionTx == param.TransactionTx && deposit.UserId == userId && param.SubscriptionId == deposit.SubscriptionId {
		err := d.db.Model(models.Deposit{}).Where("id = ?", param.Id).Update("status", param.NewStatus).Error
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New(fmt.Sprintf("consumer id :%s not valid, other col not confirm", param.Id))
}
