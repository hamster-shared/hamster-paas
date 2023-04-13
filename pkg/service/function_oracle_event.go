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
	"gorm.io/gorm"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/consts"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/rpc/eth"
	"hamster-paas/pkg/service/contract"
	"hamster-paas/pkg/utils/logger"
	"log"
	"time"
)

type FunctionOracleEventService struct {
	functionOracleContractAddress common.Address
	client                        *ethclient.Client
	db                            *gorm.DB
}

func NewFunctionOracleEventService(functionOracleContractAddress string, client *ethclient.Client, db *gorm.DB) *FunctionOracleEventService {
	return &FunctionOracleEventService{
		functionOracleContractAddress: common.HexToAddress(functionOracleContractAddress),
		client:                        client,
		db:                            db,
	}
}

func (f *FunctionOracleEventService) FunctionOracleListen() {
	chainLinkPoolService, _ := application.GetBean[*PoolService]("chainLinkPoolService")
	chainLinkPoolService.Submit(func() {
		f.oracleRequestListen()
	})
}

func (f *FunctionOracleEventService) oracleRequestListen() {
	// 定义查询过滤器
	query := ethereum.FilterQuery{
		Addresses: []common.Address{f.functionOracleContractAddress},
		Topics: [][]common.Hash{
			{
				crypto.Keccak256Hash([]byte("OracleRequest(bytes32,address,address,uint64,address,bytes)")),
			},
		},
	}
	// 创建订阅
	logs := make(chan types.Log)
	sub, err := f.client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		logger.Errorf("create Mumbai subscribe Oracle Request failed: %s", err)
	}

	contractFilter, err := contract.NewContract(f.functionOracleContractAddress, f.client)
	if err != nil {
		logger.Errorf("create Oracle Request failed: %s", err)
	}
	// 监听订阅事件
	for {
		select {
		case err := <-sub.Err():
			logger.Errorf("subscribe Oracle Request event failed: %s", err)
		case vLog := <-logs:
			logger.Info("start watch Oracle Request event")
			data, err := contractFilter.ParseOracleRequest(vLog)
			if err == nil {
				ethStr := hex.EncodeToString(data.RequestId[:])
				var subscriptionData models.Subscription
				err = f.db.Model(models.Subscription{}).Where("chain_subscription_id=? and network=?", data.SubscriptionId, eth.MUMBAI_TESTNET).First(&subscriptionData).Error
				if err == nil {
					var execData models.RequestExecute
					err = f.db.Model(models.RequestExecute{}).Where("transaction_tx=?", vLog.TxHash.Hex()).First(&execData).Error
					if err == nil {
						execData.RequestId = fmt.Sprintf("0x%s", ethStr)
						f.db.Save(&execData)
					} else {
						log.Println(ethStr)
						execData.RequestId = fmt.Sprintf("0x%s", ethStr)
						execData.SubscriptionId = int64(subscriptionData.Id)
						execData.UserId = subscriptionData.UserId
						execData.Created = time.Now()
						execData.TransactionTx = vLog.TxHash.Hex()
						execData.Status = consts.SUCCESS
						execData.ConsumerAddress = data.RequestingContract.Hex()
						execData.RequestName = ""
						execData.Args = ""
						execData.Secretsloction = 0
						execData.SecretUrl = ""
						f.db.Create(&execData)
					}
				}
			} else {
				logger.Errorf("parse OracleRequest data failed: %s", err)
			}
		}
	}
}
