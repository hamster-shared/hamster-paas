package service

import (
	"errors"
	"fmt"
	"hamster-paas/pkg/consts"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/models/vo"
	"hamster-paas/pkg/rpc/eth"
	"hamster-paas/pkg/utils/logger"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
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
func (s *ChainLinkSubscriptionService) CreateSubscription(subscription models.Subscription, poolService PoolService, network models.NetworkType) (int64, error) {
	err := s.db.Create(&subscription).Error
	if err != nil {
		return -1, err
	}
	// 异步 Tx 判断，更改 Status
	poolService.Submit(func() {
		checkAndChangeSubscriptionStatus(network, subscription, s.db)
	})
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
	var list []models.Subscription
	err := s.db.Model(models.Subscription{}).Where("user_id = ? AND status = ?", userId, consts.SUCCESS).Find(&list).Error
	if err != nil {
		return nil, err
	}
	var vo_list []vo.ChainLinkValidSubscriptionVo
	for _, v := range list {
		var s_vo vo.ChainLinkValidSubscriptionVo
		s_vo.Id = v.Id
		s_vo.ChainSubscriptionId = v.ChainSubscriptionId
		s_vo.Admin = v.Admin
		s_vo.Name = v.Name
		s_vo.Created = v.Created
		s_vo.TransactionTx = v.TransactionTx
		s_vo.Admin = v.Admin
		s_vo.Status = v.Status
		s_vo.ChainAndNetwork = v.Chain + " " + v.Network
		s_vo.NetworkId, s_vo.NetworkUrl = models.GetNetworkIdAndUrl(v.Network)
		vo_list = append(vo_list, s_vo)
	}
	return vo_list, nil
}

func (s *ChainLinkSubscriptionService) ChangeSubscriptionStatus(param vo.ChainLinkSubscriptionUpdateParam, userId uint64) error {
	//获取 id 对应的记录
	var subscription models.Subscription
	err := s.db.Model(models.Subscription{}).Where("id = ?", param.Id).First(&subscription).Error
	if err != nil {
		return err
	}
	// 如果已经是成功状态就不做操作
	if subscription.Status == param.NewStatus {
		return nil
	}
	// 判断该 consumer 是否是符合要求
	if subscription.TransactionTx == param.TransactionTx && subscription.UserId == userId && param.Chain == subscription.Chain && param.Network == subscription.Network {
		err = s.db.Model(models.Subscription{}).Where("id = ?", param.Id).Updates(map[string]interface{}{"chain_Subscription_id": param.ChainSubscriptionId, "status": param.NewStatus}).Error
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New(fmt.Sprintf("subscription id :%d not valid, other col not confirm", param.Id))
}

func (s *ChainLinkSubscriptionService) GetUserSubscriptionBalanceAll(userId int64) (*[]SubscriptionBalance, error) {
	var subscriptions []models.Subscription
	err := s.db.Model(models.Subscription{}).Where("user_id = ? AND status = ?", userId, consts.SUCCESS).Find(&subscriptions).Error
	if err != nil {
		return nil, err
	}
	var subscriptionBalances []SubscriptionBalance
	for _, subscription := range subscriptions {
		chainType, _ := models.ParseChainType(subscription.Chain)
		networkType, _ := models.ParseNetworkType(subscription.Network)
		if chainType != models.Polygon || networkType != models.TestnetMumbai {
			subscriptionBalances = append(subscriptionBalances, otherChainOrNetworkNotSupportedYet(subscription))
			continue
		}

		balance, err := GetMumbaiSubscriptionBalance(uint64(subscription.ChainSubscriptionId))
		if err != nil {
			logger.Errorf("GetMumbaiSubscriptionBalance error: %s", err)
			return nil, err
		}
		subscriptionBalance := SubscriptionBalance{
			Chain:               subscription.Chain,
			Network:             subscription.Network,
			SubscriptionId:      subscription.Id,
			ChainSubscriptionId: uint64(subscription.ChainSubscriptionId),
			Balance:             balance,
		}
		subscriptionBalances = append(subscriptionBalances, subscriptionBalance)
	}
	return &subscriptionBalances, nil
}

func otherChainOrNetworkNotSupportedYet(subscription models.Subscription) SubscriptionBalance {
	return SubscriptionBalance{
		Chain:               subscription.Chain,
		Network:             subscription.Network,
		SubscriptionId:      subscription.Id,
		ChainSubscriptionId: uint64(subscription.ChainSubscriptionId),
		Balance:             0,
		Message:             "The current chain or network does not support it",
	}
}

func (s *ChainLinkSubscriptionService) GetUserSubscriptionBalanceById(userId int64, subscriptionId uint64) (*SubscriptionBalance, error) {
	var subscription models.Subscription
	err := s.db.Model(models.Subscription{}).Where("user_id = ? AND id = ? AND status = ?", userId, subscriptionId, consts.SUCCESS).First(&subscription).Error
	if err != nil {
		return nil, err
	}
	chainType, _ := models.ParseChainType(subscription.Chain)
	networkType, _ := models.ParseNetworkType(subscription.Network)
	if chainType != models.Polygon || networkType != models.TestnetMumbai {
		result := otherChainOrNetworkNotSupportedYet(subscription)
		return &result, nil
	}
	balance, err := GetMumbaiSubscriptionBalance(uint64(subscription.ChainSubscriptionId))
	if err != nil {
		return nil, err
	}
	return &SubscriptionBalance{
		Chain:               subscription.Chain,
		Network:             subscription.Network,
		SubscriptionId:      subscription.Id,
		ChainSubscriptionId: uint64(subscription.ChainSubscriptionId),
		Balance:             balance,
	}, nil
}

type SubscriptionBalance struct {
	Chain               string `json:"chain"`
	Network             string `json:"network"`
	SubscriptionId      uint   `json:"subscriptionId"`
	ChainSubscriptionId uint64 `json:"chainSubscriptionId"`
	Balance             uint64 `json:"balance"`
	Message             string `json:"message"`
}

// 用于检查 tx 的状态，并且修改 subscription 的 status
func checkAndChangeSubscriptionStatus(network models.NetworkType, subscription models.Subscription, db *gorm.DB) {
	client := eth.GetChainClient(network.NetworkType())
	if client == nil {
		return
	}
	times := 0
	needFalid := false
	for {
		if times == 90 {
			needFalid = true
			break
		}
		time.Sleep(time.Second * 20)
		times++
		// 拿到数据库中状态，判断是否要主动结束轮询
		var s_ models.Subscription
		db.Model(models.Subscription{}).Where("id = ?", subscription.Id).First(&s_)
		// status == Success，主动结束轮询
		if s_.Status == consts.SUCCESS {
			break
		}
		// 获取 tx 状态
		txStatus, err := eth.GetTxStatus(subscription.TransactionTx, network.NetworkType(), client)
		if err != nil {
			continue
		}
		if txStatus == 1 {
			// 修改状态为成功
			logger.Infof("Create Subscription : Tx Success, change Subscription id: %d status to success", subscription.Id)
			db.Model(models.Subscription{}).Where("id = ?", subscription.Id).Update("status", consts.SUCCESS)
			break
		} else if txStatus == 0 {
			// 修改状态为失败
			logger.Infof("Create Subscription : Tx failed, change Subscription id: %d status to failed", subscription.Id)
			db.Model(models.Subscription{}).Where("id = ?", subscription.Id).Update("status", consts.FAILED)
			break
		}
	}
	if needFalid {
		// 更新状态为失败
		logger.Infof("Create Subscription : Query timeout, change Subscription id: %d status to failed", subscription.Id)
		db.Model(models.Subscription{}).Where("id = ?", subscription.Id).Update("status", consts.FAILED)
	}
}
