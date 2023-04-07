package service

import (
	"errors"
	"fmt"
	"hamster-paas/pkg/consts"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/models/vo"
	"hamster-paas/pkg/rpc/aline"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ChainLinkRequestService struct {
	db *gorm.DB
}

func NewChainLinkRequestService(db *gorm.DB) *ChainLinkRequestService {
	return &ChainLinkRequestService{
		db: db,
	}
}

func (r *ChainLinkRequestService) RequestList(page, size int, userId int64) (*vo.ChainLinkRequestPage, error) {
	var total int64
	var chainLinkRequestPage vo.ChainLinkRequestPage
	var chainLinkRequestList []models.Request
	var chainLinkRequestVoList []vo.ChainLinkRequestVo
	tx := r.db.Model(models.Request{}).Where("user_id = ?", userId)
	result := tx.Order("created DESC").Offset((page - 1) * size).Limit(size).Find(&chainLinkRequestList).Offset(-1).Limit(-1).Count(&total)
	if result.Error != nil {
		return &chainLinkRequestPage, result.Error
	}
	copier.Copy(&chainLinkRequestVoList, &chainLinkRequestList)
	chainLinkRequestPage.Data = chainLinkRequestVoList
	chainLinkRequestPage.Total = total
	chainLinkRequestPage.Page = page
	chainLinkRequestPage.PageSize = size
	return &chainLinkRequestPage, nil
}

func (r *ChainLinkRequestService) SaveChainLinkRequest(saveData vo.ChainLinkRequest) error {
	var chainLinkRequest models.Request
	err := r.db.Where("name = ? and user_id = ? ", saveData.Name, saveData.UserId).First(&chainLinkRequest).Error
	if err == gorm.ErrRecordNotFound {
		chainLinkRequest.UserId = saveData.UserId
		chainLinkRequest.Name = saveData.Name
		chainLinkRequest.Script = saveData.Script
		chainLinkRequest.Created = time.Now()
		err = r.db.Create(&chainLinkRequest).Error
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New(fmt.Sprintf("chainlink request:%s already exists", saveData.Name))
}

func (r *ChainLinkRequestService) UpdateChainLinkRequest(id int64, updateData vo.ChainLinkRequest) error {
	var chainLinkRequest models.Request
	err := r.db.Where("name = ? and user_id = ?", updateData.Name, updateData.UserId).First(&chainLinkRequest).Error
	if err == gorm.ErrRecordNotFound {
		result := r.db.Model(chainLinkRequest).Where("id = ?", id).Updates(models.Request{Name: updateData.Name, Script: updateData.Script})
		if result.Error != nil {
			return result.Error
		}
		return nil
	}
	return errors.New(fmt.Sprintf("chainlink request :%s already exists", updateData.Name))
}

func (r *ChainLinkRequestService) ChainLinkRequestTemplateList() ([]vo.RequestTemplateVo, error) {
	var templates []models.RequestTemplate
	var templateVoList []vo.RequestTemplateVo
	err := r.db.Model(models.RequestTemplate{}).Find(&templates).Error
	if err != nil {
		return templateVoList, err
	}
	copier.Copy(&templateVoList, templates)
	return templateVoList, nil
}

func (r *ChainLinkRequestService) GetRequestTemplateScript(id int64) (string, error) {
	var template models.RequestTemplate
	err := r.db.Model(models.RequestTemplate{}).Where("id = ? ", id).First(&template).Error
	if err != nil {
		return "", err
	}
	return template.Script, nil
}

func (r *ChainLinkRequestService) ChainLinkExpenseList(subscriptionId, page, size int, userId int64, requestName string) (*vo.ChainLinkExpensePage, error) {
	var total int64
	var chainLinkExpensePage vo.ChainLinkExpensePage
	var chainLinkExpenseList []models.RequestExecute
	var chainLinkExpenseVoList []vo.ChainLinkExpenseVo
	tx := r.db.Model(models.RequestExecute{}).Where("user_id = ? and subscription_id = ?", userId, subscriptionId)
	if requestName != "" {
		tx = tx.Where("request_name like ? ", "%"+requestName+"%")
	}
	result := tx.Order("created DESC").Offset((page - 1) * size).Limit(size).Find(&chainLinkExpenseList).Offset(-1).Limit(-1).Count(&total)
	if result.Error != nil {
		return &chainLinkExpensePage, result.Error
	}
	copier.Copy(&chainLinkExpenseVoList, &chainLinkExpenseList)
	chainLinkExpensePage.Data = chainLinkExpenseVoList
	chainLinkExpensePage.Total = total
	chainLinkExpensePage.Page = page
	chainLinkExpensePage.PageSize = size
	return &chainLinkExpensePage, nil
}

func (r *ChainLinkRequestService) SaveChainLinkRequestExec(saveData vo.ChainLinkRequestExecParam, userId uint64) error {
	//todo 链上校验
	var requestExec models.RequestExecute
	copier.Copy(&requestExec, &saveData)
	requestExec.Created = time.Now()
	requestExec.UserId = userId
	requestExec.Status = consts.PENDING
	err := r.db.Create(&requestExec).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ChainLinkRequestService) Overview(user aline.User) (error, error) {

	return nil, nil
}
