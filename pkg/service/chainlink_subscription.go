package service

import (
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

func (s *ChainLinkSubscriptionService) CreateSubscription(subscription models.Subscription) {
	s.db.Create(&subscription)
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

//// GetSINAByUserId ge t Subscription id,name,balance by user_id
//func (s *ChainLinkSubscriptionService) GetSINAByUserId(UserId uint) []*SINA {
//	var sinas []*SINA
//	s.db.Table("t_cl_subscription").
//		Select("subscription_id", "name", "balance").
//		Where("user_id = ?", UserId).
//		Scan(&sinas)
//
//	return sinas
//}
