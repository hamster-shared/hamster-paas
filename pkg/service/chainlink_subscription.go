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
)

type ChainLinkSubscriptionService struct {
	db *gorm.DB
}

func NewChainLinkSubscriptionService(db *gorm.DB) *ChainLinkSubscriptionService {
	return &ChainLinkSubscriptionService{
		db: db,
	}
}

// CreateSubscription create subscription
// * param subscription: new Subscription need to save in db.
// * error when subscription already exit.
func (s *ChainLinkSubscriptionService) CreateSubscription(subscription models.Subscription, poolService PoolService, network eth.EthNetwork) (int64, error) {
	err := s.db.Create(&subscription).Error
	if err != nil {
		return -1, err
	}
	//poolService.Submit(func() {
	//	client, _ := eth.NewRPCEthereumProxy(eth.NetMap[network])
	//	times := 0
	//	needFalid := false
	//	for {
	//		if times == 300 {
	//			needFalid = true
	//			break
	//		}
	//		re, _ := client.TransactionReceipt(subscription.TransactionTx)
	//		if re.Status == 1 {
	//			// 修改状态为成功
	//			logger.Infof("Create Subscription : Tx Success, change subscription id: %d status to success", subscription.Id)
	//			s.db.Model(models.Subscription{}).Where("id = ?", subscription.Id).Update("status", consts.SUCCESS)
	//			break
	//		} else if re.Status == 0 {
	//			// 修改状态为失败
	//			logger.Infof("Create Subscription : Tx failed, change subscription id: %d status to failed", subscription.Id)
	//			s.db.Model(models.Subscription{}).Where("id = ?", subscription.Id).Update("status", consts.FAILED)
	//			break
	//		}
	//		time.Sleep(time.Minute * 1)
	//		times++
	//	}
	//	if needFalid {
	//		// 更新状态为失败
	//		logger.Infof("Create Subscription : Query timeout, change subscription id: %d status to failed", subscription.Id)
	//		s.db.Model(models.Subscription{}).Where("id = ?", subscription.Id).Update("status", consts.FAILED)
	//	}
	//
	//})

	return int64(subscription.Id), nil
}

// GetSubscriptionOverview get subscription overview(subscription nums, consumer nums, balances)
// * param userId: user id.
// * param network: Test or Main
// * return overview.
func (s *ChainLinkSubscriptionService) GetSubscriptionOverview(userId uint, network string) (*vo.ChainLinkSubscriptionOverview, error) {
	var vo *vo.ChainLinkSubscriptionOverview
	sql := "select COUNT(*) as total_subscription, SUM(consumers) as total_consumers from t_cl_subscription where user_id = ? AND network LIKE ? AND status = ? "
	like_ := "%" + network
	if err := s.db.Raw(sql, userId, like_, consts.SUCCESS).Scan(&vo).Error; err != nil {
		return nil, err
	}
	return vo, nil
}

// GetSINAByUserId ge t Subscription id,name,balance by user_id
func (s *ChainLinkSubscriptionService) GetSINAByUserId(UserId uint) []*vo.ChainLinkSINA {
	var sinas []*vo.ChainLinkSINA
	s.db.Model(models.Subscription{}).
		Select("subscription_id", "name", "balance").
		Where("user_id = ?", UserId).
		Scan(&sinas)

	return sinas
}

// UpdateConsumerNums update consumer for subscription
// param subscriptionId: which subscription
// param consumerNums: the subscription new consumer nums
func (s *ChainLinkSubscriptionService) UpdateConsumerNums(subscriptionId uint, newConsumerNums int64) error {
	s.db.Model(models.Subscription{}).Where("id = ?", subscriptionId).Update("consumers", newConsumerNums)
	return nil
}

func (s *ChainLinkSubscriptionService) GetSubscriptionById(id int) (*models.Subscription, error) {
	var subscription *models.Subscription
	if err := s.db.Model(models.Subscription{}).Where("id = ? AND status = ?", id, consts.SUCCESS).First(&subscription).Error; err != nil {
		return nil, err
	}
	return subscription, nil
}

// SubscriptionList  query subscription list
func (s *ChainLinkSubscriptionService) SubscriptionList(chain, network string, page, size int, userId int64) (*vo.ChainLinkSubscriptionPage, error) {
	var total int64
	var chainLinkSubscriptionPage vo.ChainLinkSubscriptionPage
	var chainLinkSubscriptionList []models.Subscription
	var chainLinkSubscriptionVoList []vo.ChainLinkSubscriptionVo
	tx := s.db.Model(models.Subscription{}).Where("user_id = ?", userId)
	if network != "" && chain != "" {
		tx = tx.Where("network = ? and chain = ?", network, chain)
	}
	result := tx.Order("created DESC").Offset((page - 1) * size).Limit(size).Find(&chainLinkSubscriptionList).Offset(-1).Limit(-1).Count(&total)
	if result.Error != nil {
		return &chainLinkSubscriptionPage, result.Error
	}
	copier.Copy(&chainLinkSubscriptionVoList, &chainLinkSubscriptionList)
	chainLinkSubscriptionPage.Data = chainLinkSubscriptionVoList
	chainLinkSubscriptionPage.Total = total
	chainLinkSubscriptionPage.Page = page
	chainLinkSubscriptionPage.PageSize = size
	return &chainLinkSubscriptionPage, nil
}

// SubscriptionDetail query subscription detail by id
func (s *ChainLinkSubscriptionService) SubscriptionDetail(id int64) (vo.ChainLinkSubscriptionVo, error) {
	var subscriptionData models.Subscription
	var vo vo.ChainLinkSubscriptionVo
	err := s.db.Model(models.Subscription{}).Where("id = ?", id).First(&subscriptionData).Error
	if err != nil {
		return vo, err
	}
	copier.Copy(&vo, &subscriptionData)
	return vo, nil
}

func (s *ChainLinkSubscriptionService) GetValidSubscription(userId int64) ([]vo.ChainLinkValidSubscriptionVo, error) {
	var list []vo.ChainLinkValidSubscriptionVo
	err := s.db.Model(models.Subscription{}).Where("user_id = ? AND status = ?", userId, consts.SUCCESS).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *ChainLinkSubscriptionService) ChangeSubscriptionStatus(param vo.ChainLinkSubscriptionUpdateParam, userId uint64) error {
	//获取id对应的记录
	var subscription models.Subscription
	err := s.db.Model(models.Subscription{}).Where("id = ?", param.Id).First(&subscription).Error
	if err != nil {
		return err
	}
	// 如果已经是成功状态就不做操作
	if subscription.Status == consts.SUCCESS {
		return nil
	}
	// 判断该consumer是否是符合要求
	if subscription.TransactionTx == param.TransactionTx && subscription.UserId == userId && param.Chain == subscription.Chain && param.Network == subscription.Network {
		fmt.Println(param.ChainSubscriptionId)
		err = s.db.Model(models.Subscription{}).Where("id = ?", param.Id).Updates(map[string]interface{}{"chain_Subscription_id": param.ChainSubscriptionId, "status": param.NewStatus}).Error
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New(fmt.Sprintf("subscription id :%s not valid, other col not confirm", param.Id))
}
