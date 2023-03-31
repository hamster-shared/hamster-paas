package service

import (
	"errors"
	"fmt"
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
