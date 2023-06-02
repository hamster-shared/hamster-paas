package service

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"gorm.io/gorm"
	"hamster-paas/pkg/consts"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/rpc/eth"
	"hamster-paas/pkg/service/contract"
	"hamster-paas/pkg/utils/logger"
	"math"
	"math/big"
	"strings"
	"time"
)

type BillingContractEventService struct {
	billingContractAddress common.Address
	client                 *ethclient.Client
	db                     *gorm.DB
	network                eth.EthNetwork
}

func NewBillingContractEventService(billingContractAddress string, client *ethclient.Client, db *gorm.DB, network eth.EthNetwork) *BillingContractEventService {
	return &BillingContractEventService{
		billingContractAddress: common.HexToAddress(billingContractAddress),
		client:                 client,
		db:                     db,
		network:                network,
	}
}

func (b *BillingContractEventService) BillingRegistryListen(errorChan chan error, chanTwo chan error) {
	go func() {
		err := b.billingEndListen()
		errorChan <- err
	}()
	go func() {
		err := b.subscriptionFundedListen()
		chanTwo <- err
	}()
}

func (b *BillingContractEventService) billingEndListen() error {
	// 定义查询过滤器
	query := ethereum.FilterQuery{
		Addresses: []common.Address{b.billingContractAddress},
		Topics: [][]common.Hash{
			{
				crypto.Keccak256Hash([]byte("BillingEnd(bytes32,uint64,uint96,uint96,uint96,bool)")),
			},
		},
	}
	// 创建订阅
	logs := make(chan types.Log)
	sub, err := b.client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		logger.Errorf("create Mumbai subscription billing end failed: %s", err)
		return err
	}

	contractFilter, err := contract.NewBillingRegistryFilterer(b.billingContractAddress, b.client)
	if err != nil {
		logger.Errorf("create billing registry filter failed: %s", err)
		return err
	}
	// 监听订阅事件
	for {
		select {
		case err := <-sub.Err():
			logger.Errorf("subscription billing end event failed: %s", err)
			return err
		case vLog := <-logs:
			logger.Info("start watch billing end event:")
			data, err := contractFilter.ParseBillingEnd(vLog)
			if err == nil {
				b.handleBillingEndData(data)
			} else {
				logger.Errorf("parse billing end data failed: %s", err)
				return err
			}
		}
	}
}

func (b *BillingContractEventService) handleBillingEndData(data *contract.BillingRegistryBillingEnd) {
	ethStr := hex.EncodeToString(data.RequestId[:])
	var subscriptionData models.Subscription
	err := b.db.Model(models.Subscription{}).Where("chain_subscription_id=? and network=?", data.SubscriptionId, b.network).First(&subscriptionData).Error
	if err == nil {
		amount, _ := weiToEth(data.TotalCost).Float64()
		if subscriptionData.Balance > amount {
			subscriptionData.Balance = subscriptionData.Balance - amount
			b.db.Save(&subscriptionData)
		}
		var execData models.RequestExecute
		err = b.db.Model(models.RequestExecute{}).Where("request_id=?", fmt.Sprintf("0x%s", ethStr)).First(&execData).Error
		if err == nil {
			execData.Amount = amount
			b.db.Save(&execData)
		}
	}
}

func (b *BillingContractEventService) subscriptionFundedListen() error {
	// 定义查询过滤器
	query := ethereum.FilterQuery{
		Addresses: []common.Address{b.billingContractAddress},
		Topics: [][]common.Hash{
			{
				crypto.Keccak256Hash([]byte("SubscriptionFunded(uint64,uint256,uint256)")),
			},
		},
	}
	// 创建订阅
	logs := make(chan types.Log)
	sub, err := b.client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		logger.Errorf("create Mumbai subscription Subscription Funded failed: %s", err)
		return err
	}

	contractFilter, err := contract.NewBillingRegistryFilterer(b.billingContractAddress, b.client)
	if err != nil {
		logger.Errorf("crate Subscription Funded filter failed: %s", err)
		return err
	}
	// 监听订阅事件
	for {
		select {
		case err := <-sub.Err():
			logger.Errorf("订阅 Subscription Funded event出错: %s", err)
			return err
		case vLog := <-logs:
			logger.Info("watch Subscription Funded event")
			data, err := contractFilter.ParseSubscriptionFunded(vLog)
			if err == nil {
				tx, isPending, _ := b.client.TransactionByHash(context.Background(), vLog.TxHash)
				if !isPending {
					b.handleSubscriptionFundedData(data, tx, vLog)
				}
			} else {
				logger.Errorf("parse SubscriptionFunded data failed: %s", err)
				return err
			}
		}
	}
}

func (b *BillingContractEventService) handleSubscriptionFundedData(data *contract.BillingRegistrySubscriptionFunded, tx *types.Transaction, vLog types.Log) {
	var subscriptionData models.Subscription
	logger.Info("Start handle data--------------------")
	err := b.db.Model(models.Subscription{}).Where("chain_subscription_id=? and network=?", data.SubscriptionId, b.network).First(&subscriptionData).Error
	if err == nil {
		amount, _ := weiToEth(data.NewBalance).Float64()
		logger.Info("+++++++++++++++++")
		logger.Info(b.network)
		logger.Info(data.SubscriptionId)
		logger.Info(amount)
		logger.Info("+++++++++++++++++")
		subscriptionData.Balance = amount
		b.db.Save(&subscriptionData)
		signer := types.NewEIP155Signer(tx.ChainId())
		fromAddress, err := signer.Sender(tx)
		if err == nil {
			var status string
			var errorMessage string
			re, err := eth.GetTxStatus(tx.Hash().Hex(), b.network, b.client)
			if err != nil {
				logger.Errorf("get tx status error: %s", err)
				status = consts.FAILED
				errorMessage = "get tx failed"
			}
			if re.Status == 1 {
				status = consts.SUCCESS
			}
			if re.Status == 0 {
				if len(re.Logs) > 0 {
					event := &types.Log{}
					err = rlp.DecodeBytes(re.Logs[0].Data, event)
					if err != nil {
						errorMessage = err.Error()
					}
					errMsg := string(event.Data)
					errMsg = strings.Trim(errMsg, "\x00")
					errorMessage = errMsg
				} else {
					errorMessage = "Transaction failed without error information"
				}
			}
			var depositData models.Deposit
			depositData.SubscriptionId = int64(subscriptionData.Id)
			depositData.Status = status
			depositData.TransactionTx = vLog.TxHash.Hex()
			depositData.UserId = subscriptionData.UserId
			depositData.Created = time.Now()
			depositData.Amount = amount
			depositData.Address = fromAddress.Hex()
			depositData.ErrorMessage = errorMessage
			b.db.Create(&depositData)
		}
	}
}

func weiToEth(wei *big.Int) *big.Float {
	eth := new(big.Float)
	eth.SetString(wei.String())
	return new(big.Float).Quo(eth, big.NewFloat(math.Pow10(18)))
}
