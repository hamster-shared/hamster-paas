package service

import (
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/models/vo"
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
func (d *ChainLinkDepositService) DepositList(page, size int, userId int64) (*vo.ChainLinkDepositPage, error) {
	var total int64
	var chainLinkDepositPage vo.ChainLinkDepositPage
	var chainLinkDepositList []models.Deposit
	var chainLinkDepositVoList []vo.ChainLinkDepositVo
	tx := d.db.Model(models.Deposit{}).Where("user_id = ?", userId)
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