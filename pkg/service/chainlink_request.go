package service

import (
	"errors"
	"fmt"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/consts"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/models/vo"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/rpc/eth"
	"hamster-paas/pkg/utils/logger"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"

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
	tx := r.db.Model(models.RequestExecute{}).Where("user_id = ? and subscription_id = ? and status = ? ", userId, subscriptionId, consts.SUCCESS)
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

func (r *ChainLinkRequestService) SaveChainLinkRequestExec(saveData vo.ChainLinkRequestExecParam, user aline.User) (int64, error) {
	var requestExec models.RequestExecute
	copier.Copy(&requestExec, &saveData)
	requestExec.Created = time.Now()
	requestExec.UserId = uint64(user.Id)
	requestExec.Status = consts.PENDING
	err := r.db.Create(&requestExec).Error
	if err != nil {
		return 0, err
	}
	client := eth.NewEthereumProxyFactory().GetClient(eth.EthNetwork(saveData.Network))
	chainLinkPoolService, err := application.GetBean[*PoolService]("chainLinkPoolService")
	if err != nil {
		logger.Error(fmt.Sprintf("get pool service failed:%s", err.Error()))
		return requestExec.Id, nil
	}
	statusFun := func() {
		watchExecStatus(requestExec, r.db, client)
		watchRequest(saveData.ConsumerAddress, saveData.RequestId, user.UserEmail, client)
	}
	chainLinkPoolService.Submit(statusFun)
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

func (r *ChainLinkRequestService) Overview(user aline.User, networkType models.NetworkType) (*models.ApiResponseOverview, error) {
	sqlQuery := `SELECT *
FROM t_cl_subscription
JOIN t_cl_oracle_request_event
ON t_cl_subscription.chain_subscription_id = t_cl_oracle_request_event.subscription_id
WHERE t_cl_subscription.chain_subscription_id IN (
    SELECT chain_subscription_id FROM t_cl_subscription WHERE user_id = ?
)
AND t_cl_subscription.chain = ?
AND t_cl_subscription.network = ?
`
	var result []models.OracleRequestEventAndName
	err := r.db.Raw(sqlQuery, user.Id, "ethereum", networkType.StringWithSpace()).Scan(&result).Error
	if err != nil {
		logger.Errorf("chain link oracle request overview error: %s", err)
		return nil, err
	}

	var apiResponseOverview models.ApiResponseOverview
	apiResponseOverview.Network = networkType.StringWithSpace()
	// 首先过滤出种类
	for _, v := range result {
		if !contains(apiResponseOverview.LegendData, v.Name) {
			apiResponseOverview.LegendData = append(apiResponseOverview.LegendData, v.Name)
		}
	}
	for i := 0; i < 7; i++ {
		apiResponseOverview.XaxisData = append(apiResponseOverview.XaxisData, time.Now().AddDate(0, 0, -i).Format("2006-01-02"))
	}
	apiResponseOverview.XaxisData = reverseString(apiResponseOverview.XaxisData)

	// 根据时间过滤，只保留最近 7 天的请求事件，每天一个，最后成一个数组，最终只要出现的次数而已
	for _, v := range apiResponseOverview.LegendData {
		var serie models.Serie
		serie.Name = v
		for _, x := range apiResponseOverview.XaxisData {
			var count int
			for _, r := range result {
				if r.Name == v && strings.Contains(r.CreatedAt.Format("2006-01-02"), x) {
					count++
				}
			}
			serie.Data = append(serie.Data, count)
		}
		apiResponseOverview.SeriesData = append(apiResponseOverview.SeriesData, serie)
	}
	return &apiResponseOverview, nil
}

// 查看某个键是否在列表里
func contains(list []string, key string) bool {
	for _, v := range list {
		if v == key {
			return true
		}
	}
	return false
}

func reverseString(in []string) []string {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
	return in
}

func watchRequest(contractAddress, requestId, email string, client eth.EthereumProxy) {
	client.WatchRequestResult(contractAddress, requestId, email)
}

func watchExecStatus(data models.RequestExecute, db *gorm.DB, client eth.EthereumProxy) {
	start := time.Now() // 记录开始时间
	for {
		receipt, err := client.TransactionReceipt(data.TransactionTx)
		if err != nil {
			data.Status = consts.FAILED
			db.Save(&data)
			break
		}
		if receipt.Status == types.ReceiptStatusFailed {
			data.Status = consts.FAILED
			db.Save(&data)
			break
		}
		if receipt.Status == types.ReceiptStatusSuccessful {
			data.Status = consts.SUCCESS
			db.Save(&data)
			break
		}
		if time.Since(start) > time.Minute*3 {
			break
		}
	}
}
