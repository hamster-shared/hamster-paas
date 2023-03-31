package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/models/vo"
	"log"
)

type ChainLinkSubscriptionService struct {
	db *gorm.DB
}

func NewChainLinkSubscriptionService(db *gorm.DB) *ChainLinkSubscriptionService {
	return &ChainLinkSubscriptionService{
		db: db,
	}
}

func (s *ChainLinkSubscriptionService) CreateSubscription(subscription models.Subscription) error {
	var s_ *models.Subscription
	// 判断subscription id是否存在
	err := s.db.Table("t_cl_subscription").Where("subscription_id = ?", subscription.SubscriptionId).First(&s_).Error
	// 不存在则创建
	if err == gorm.ErrRecordNotFound {
		err = s.db.Create(&subscription).Error
		if err != nil {
			return err
		}
		return nil
	}
	// 已存在，返回错误
	return errors.New(fmt.Sprintf("subscription :%d already exists", subscription.SubscriptionId))
}

func (s *ChainLinkSubscriptionService) GetSubscriptionOverview(userId uint, network string) (*vo.ChainLinkSubscriptionOverview, error) {

	var vo *vo.ChainLinkSubscriptionOverview

	sql := "select COUNT(*) as total_subscription, SUM(consumers) as total_consumers, SUM(balance) as total_balance from t_cl_subscription where user_id = ? AND network = ?"
	if err := s.db.Raw(sql, userId, network).Scan(&vo).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return vo, nil
}

// GetSINAByUserId ge t Subscription id,name,balance by user_id
func (s *ChainLinkSubscriptionService) GetSINAByUserId(UserId uint) []*vo.ChainLinkSINA {
	var sinas []*vo.ChainLinkSINA
	s.db.Table("t_cl_subscription").
		Select("subscription_id", "name", "balance").
		Where("user_id = ?", UserId).
		Scan(&sinas)

	return sinas
}

func (s *ChainLinkSubscriptionService) AddConsumer(subscriptionId uint, consumerNums int64) error {
	s.db.Table("t_cl_subscription").Where("subscription_id = ?", subscriptionId).Update("consumers", consumerNums)
	return nil
}

func (s *ChainLinkSubscriptionService) GetSubscriptionById(id int) (*models.Subscription, error) {
	var subscription *models.Subscription

	if err := s.db.Table("t_cl_subscription").Where("subscription_id = ?", id).First(&subscription).Error; err != nil {
		return nil, err
	}

	return subscription, nil
}

func (s *ChainLinkSubscriptionService) AddFundsForSubscription(subscriptionId int, incr float64) error {
	subscription, err := s.GetSubscriptionById(subscriptionId)
	if err != nil {
		return err
	}

	err = s.db.Table("t_cl_subscription").Where("subscription_id = ?", subscriptionId).Update("balance", subscription.Balance+incr).Error
	if err != nil {
		return err
	}

	return nil
}

// SubscriptionList  query subscription list
func (s *ChainLinkSubscriptionService) SubscriptionList(network string, page, size int, userId int64) (*vo.ChainLinkSubscriptionPage, error) {
	var total int64
	var chainLinkSubscriptionPage vo.ChainLinkSubscriptionPage
	var chainLinkSubscriptionList []models.Subscription
	var chainLinkSubscriptionVoList []vo.ChainLinkSubscriptionVo
	tx := s.db.Model(models.Subscription{}).Where("user_id = ?", userId)
	if network != "" {
		tx = tx.Where("network = ?", network)
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
