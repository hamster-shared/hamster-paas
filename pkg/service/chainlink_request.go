package service

import (
	"errors"
	"fmt"
	"hamster-paas/pkg/consts"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/models/vo"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/utils/logger"
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
		chainLinkRequest.ParamsCount = saveData.ParamsCount
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

func (r *ChainLinkRequestService) GetRequestTemplateScript(id int64) (vo.RequestTemplateDetailVo, error) {
	var template models.RequestTemplate
	var detail vo.RequestTemplateDetailVo
	err := r.db.Model(models.RequestTemplate{}).Where("id = ? ", id).First(&template).Error
	if err != nil {
		return detail, err
	}
	copier.Copy(&detail, &template)
	return detail, nil
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

func (r *ChainLinkRequestService) SaveChainLinkRequestExec(saveData vo.ChainLinkRequestExecParam, userId uint64) (int64, error) {
	//todo 链上校验
	var requestExec models.RequestExecute
	copier.Copy(&requestExec, &saveData)
	requestExec.Created = time.Now()
	requestExec.UserId = userId
	requestExec.Status = consts.PENDING
	err := r.db.Create(&requestExec).Error
	if err != nil {
		return 0, err
	}
	return requestExec.Id, nil
}

func (r *ChainLinkRequestService) UpdateChainLinkRequestById(id, userId int64, status string) error {
	var data models.RequestExecute
	err := r.db.Where("id = ? and user_id = ?", id, userId).First(&data).Error
	if err != nil {
		return err
	}
	data.Status = status
	r.db.Save(&data)
	return nil
}

func (r *ChainLinkRequestService) Overview(user aline.User) ([]models.OracleRequestEventAndName, error) {
	sqlQuery := `SELECT *
FROM t_cl_subscription
JOIN t_cl_oracle_request_event
ON t_cl_subscription.transaction_tx = t_cl_oracle_request_event.transaction_hash
WHERE t_cl_subscription.chain_subscription_id IN (
    SELECT chain_subscription_id FROM t_cl_subscription WHERE user_id = ?
)
AND t_cl_subscription.chain = ?
AND t_cl_subscription.network = ?
`
	logger.Info("sqlQuery: ", sqlQuery)

	var result []models.OracleRequestEventAndName
	err := r.db.Raw(sqlQuery, user.Id, "ethereum", "testnet-mumbai").Scan(&result).Error
	if err != nil {
		logger.Errorf("chain link oracle request overview error: %s", err)
		return nil, err
	}

	// 根据时间过滤，只保留最近 7 天的请求事件
	// var filteredRequestEvents []models.OracleRequestEvent
	// for _, requestEvent := range requestEvents {
	// 	if time.Now().Sub(requestEvent.CreatedAt) < 7*24*time.Hour {
	// 		filteredRequestEvents = append(filteredRequestEvents, requestEvent)
	// 	}
	// }
	return result, nil
}
