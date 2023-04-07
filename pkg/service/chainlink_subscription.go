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
func (s *ChainLinkSubscriptionService) CreateSubscription(subscription models.Subscription, poolService PoolService) error {
	var s_ *models.Subscription
	// 判断是否用已经成功的订阅存在
	err := s.db.Model(models.Subscription{}).Where(
		"chain_subscription_id = ? AND chain = ? AND network = ? AND (status = ? OR status = ?)",
		subscription.ChainSubscriptionId, subscription.Chain, subscription.Network, consts.PENDING, consts.SUCCESS).First(&s_).Error
	// 判断订阅是否存在
	if err == gorm.ErrRecordNotFound {
		err = s.db.Create(&subscription).Error
		if err != nil {
			return err
		}
		// 异步检查订阅交易事务是否成功
		poolService.Submit(func() {
			time.Sleep(time.Second * 20)
			s.db.Model(models.Subscription{}).Where("chain_subscription_id = ? AND chain = ? AND network = ?", subscription.ChainSubscriptionId, subscription.Chain, subscription.Network).
				Update("status", consts.SUCCESS)
			logger.Info("create subscription: %d Tx valid, change status to success", subscription.ChainSubscriptionId)
		})
		return nil
	}
	// 订阅已存在，返回错误
	return errors.New(fmt.Sprintf("chain: %s network: %s ,subscription :%d already exists, status: %s", subscription.Chain, subscription.Network, subscription.ChainSubscriptionId, s_.Status))
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
