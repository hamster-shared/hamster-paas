package service

import (
	"context"
	"fmt"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/rpc/eth"
	"hamster-paas/pkg/utils/logger"
	"time"

	"hamster-paas/pkg/service/oracle"
	oracle_proxy "hamster-paas/pkg/service/oracle/proxy"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

var EthereumSepoliaURL = "wss://eth-sepolia.g.alchemy.com/v2/Zprj6NPmzkOwmoRLGWBo74-S-NkijXeQ"
var PolygonMumbaiURL = "wss://polygon-mumbai.g.alchemy.com/v2/ag4Hb9DuuoRxhWou2mHdJrdQdc9_JFXG"
var EthereumSepoliaOracle = "0x649a2C205BE7A3d5e99206CEEFF30c794f0E31EC"
var PolygonMumbaiOracle = "0xeA6721aC65BCeD841B8ec3fc5fEdeA6141a0aDE4"
var MumbaiBillingRegistryAddress = "0xEe9Bf52E5Ea228404bB54BCFbbDa8c21131b9039"
var MumbaiFunctionOracleAddress = "0xeA6721aC65BCeD841B8ec3fc5fEdeA6141a0aDE4"
var SepoliaBillingRegistryAddress = "0x3c79f56407DCB9dc9b852D139a317246f43750Cc"
var SepoliaFunctionOracleAddress = "0x649a2C205BE7A3d5e99206CEEFF30c794f0E31EC"

var ORACLE_BILLING_REGISTRY_PROXY = "0xee9bf52e5ea228404bb54bcfbbda8c21131b9039"

func (l *OracleListener) listenOracleRequestEvent(ethUrl, contractAddressString string) error {
	// 连接到 Ethereum 节点
	client, err := ethclient.Dial(ethUrl)
	if err != nil {
		logger.Errorf("连接到链节点出错: %s", err)
		return err
	}
	logger.Infof("已连接到链节点: %s", ethUrl)
	if ethUrl == EthereumSepoliaURL {
		l.ethereumSepoliaClient = client
	} else {
		l.polygonMumbaiClient = client
	}

	// 智能合约地址
	contractAddress := common.HexToAddress(contractAddressString)

	// 定义查询过滤器
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{}, {}},
	}

	// 创建订阅
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		logger.Errorf("创建订阅出错: %s", err)
		return err
	}

	contract, err := oracle.NewOracleFilterer(contractAddress, client)
	if err != nil {
		logger.Errorf("创建合约实例出错: %s", err)
		return err
	}

	// 监听订阅事件
	for {
		select {
		case err := <-sub.Err():
			logger.Errorf("订阅出错: %s", err)
			return err
		case vLog := <-logs:
			// fmt.Println("交易哈希：", vLog.TxHash.Hex())
			// fmt.Println("合约地址：", vLog.Address.Hex())

			switch vLog.Topics[0].Hex() {
			case getOracleRequestTopic().Hex():
				o, err := contract.ParseOracleRequest(vLog)
				if err != nil {
					logger.Errorf("解析 OracleRequest 事件出错: %s", err)
					return err
				}
				// spew.Dump(o)
				var chain, network string
				if ethUrl == EthereumSepoliaURL {
					chain = "Ethereum"
					network = "Sepolia Testnet"
				} else {
					chain = "Polygon"
					network = "Mumbai Testnet"
				}
				l.saveOracleRequestEvent(o, chain, network)

			default:
				// 非 OracleRequest 事件，不关心
			}
		}
	}
}

func getOracleRequestTopic() common.Hash {
	logOracleRequestSig := []byte("OracleRequest(bytes32,address,address,uint64,address,bytes)")
	return crypto.Keccak256Hash(logOracleRequestSig)
}

type OracleListener struct {
	polygonMumbaiClient   *ethclient.Client
	ethereumSepoliaClient *ethclient.Client
	db                    *gorm.DB
}

func NewOracleListener(db *gorm.DB) *OracleListener {
	return &OracleListener{
		db: db,
	}
}

func (l *OracleListener) StartListen() {
	mumbaiOracleChan := make(chan struct{})
	sepoliaOracleChan := make(chan struct{})
	mumbaiChan := make(chan struct{})
	sepoliaChan := make(chan struct{})
	go func() {
		for {
			logger.Info("准备监听 polygon mumbai 获取 oracle request event")
			err := l.listenOracleRequestEvent(PolygonMumbaiURL, PolygonMumbaiOracle)
			if err != nil {
				logger.Errorf("监听 polygon mumbai 出错: %s", err)
				time.Sleep(5 * time.Second)
				go func() {
					mumbaiOracleChan <- struct{}{}
				}()
			}
			<-mumbaiOracleChan
			logger.Info("准备重试连接 polygon mumbai")
		}
	}()
	go func() {
		for {
			logger.Info("准备监听 ethereum sepolia 获取 oracle request event")
			err := l.listenOracleRequestEvent(EthereumSepoliaURL, EthereumSepoliaOracle)
			if err != nil {
				logger.Errorf("监听 ethereum sepolia 出错: %s", err)
				time.Sleep(5 * time.Second)
				go func() {
					sepoliaOracleChan <- struct{}{}
				}()
			}
			<-sepoliaOracleChan
			logger.Info("准备重试连接 ethereum sepolia")
		}
	}()
	go func() {
		for {
			logger.Info("准备监听 Ethereum 获取 oracle request event")
			err := l.MumbaiListen()
			if err != nil {
				logger.Errorf("监听 eth 出错：%s", err)
				time.Sleep(5 * time.Second)
				go func() {
					mumbaiChan <- struct{}{}
				}()
			}
			<-mumbaiChan
			logger.Info("准备重试连接 Mumbai")
		}
	}()
	go func() {
		for {
			logger.Info("准备监听 sepolia 网络的 event")
			err := l.SepoliaListen()
			if err != nil {
				logger.Errorf("监听 sepolia 出错：%s", err)
				time.Sleep(5 * time.Second)
				go func() {
					sepoliaChan <- struct{}{}
				}()
			}
			<-sepoliaChan
			logger.Info("准备重试连接 sepolia")
		}
	}()
}

func (l *OracleListener) saveOracleRequestEvent(r *oracle.OracleOracleRequest, chain, network string) {
	var event models.OracleRequestEvent
	event.TransactionHash = r.Raw.TxHash.Hex()
	event.RequestingContract = r.RequestingContract.Hex()
	event.RequestInitiator = r.RequestInitiator.Hex()
	event.SubscriptionId = r.SubscriptionId
	event.SubscriptionOwner = r.SubscriptionOwner.Hex()
	event.BlockNumber = r.Raw.BlockNumber
	event.TxIndex = r.Raw.TxIndex
	event.BlockHash = r.Raw.BlockHash.Hex()
	event.Index = r.Raw.Index
	event.Removed = r.Raw.Removed
	event.Chain = chain
	event.Network = network
	event.CreatedAt = time.Now()
	err := l.db.Create(&event).Error
	if err != nil {
		logger.Errorf("保存 oracle request event 失败: %s", err)
	} else {
		logger.Infof("保存 oracle request event success: transaction hash: %s", event.TransactionHash)
	}
}

func (l *OracleListener) GetFund(subscriptionId uint64) (uint64, error) {
	if l.polygonMumbaiClient == nil {
		return 0, fmt.Errorf("eth client is nil")
	}
	contractAddress := common.HexToAddress(ORACLE_BILLING_REGISTRY_PROXY)
	caller, err := oracle_proxy.NewOracleProxyCaller(contractAddress, l.polygonMumbaiClient)
	if err != nil {
		return 0, err
	}
	result, err := caller.GetSubscription(&bind.CallOpts{}, subscriptionId)
	if err != nil {
		return 0, err
	}
	return result.Balance.Uint64(), nil
}

func GetMumbaiSubscriptionBalance(subscriptionId uint64) (uint64, error) {
	oracleListener, err := application.GetBean[*OracleListener]("oracleListener")
	if err != nil {
		return 0, err
	}
	return oracleListener.GetFund(subscriptionId)
}

func (l *OracleListener) MumbaiListen() error {
	client, err := ethclient.Dial(eth.NetMap[eth.MUMBAI_TESTNET])
	if err != nil {
		logger.Errorf("connect Mumbai node failed: %s", err)
		return err
	}
	logger.Info("connected Mumbai node")
	billingRegistryService := NewBillingContractEventService(MumbaiBillingRegistryAddress, client, l.db, eth.MUMBAI_TESTNET)
	errChan, errChan2, errChan3, errChan4 := make(chan error), make(chan error), make(chan error), make(chan error)
	billingRegistryService.BillingRegistryListen(errChan, errChan2)
	functionOracleService := NewFunctionOracleEventService(MumbaiFunctionOracleAddress, client, l.db, eth.MUMBAI_TESTNET)
	functionOracleService.FunctionOracleListen(errChan3)
	functionConsumerService := NewFunctionConsumerEventService(client, l.db, eth.MUMBAI_TESTNET)
	functionConsumerService.FunctionConsumerListen(errChan4)
	for {
		select {
		case err1 := <-errChan:
			logger.Errorf("mubai watch event error, error1 is : %s", err1)
			return err1
		case err2 := <-errChan2:
			logger.Errorf("mubai watch event error, error2 is : %s", err2)
			return err2
		case err3 := <-errChan3:
			logger.Errorf("mubai watch event error, error3 is : %s", err3)
			return err3
		case err4 := <-errChan4:
			logger.Errorf("mubai watch event error, error4 is : %s", err4)
			return err4
		}
	}
	return nil
}

func (l *OracleListener) SepoliaListen() error {
	client, err := ethclient.Dial(eth.NetMap[eth.SEPOLIA_TESTNET])
	if err != nil {
		logger.Errorf("connect Sepolia node failed: %s", err)
		return err
	}
	logger.Info("connected Sepolia node")
	billingRegistryService := NewBillingContractEventService(SepoliaBillingRegistryAddress, client, l.db, eth.SEPOLIA_TESTNET)
	errChan, errChan2, errChan3, errChan4 := make(chan error), make(chan error), make(chan error), make(chan error)
	billingRegistryService.BillingRegistryListen(errChan, errChan2)
	functionOracleService := NewFunctionOracleEventService(SepoliaFunctionOracleAddress, client, l.db, eth.SEPOLIA_TESTNET)
	functionOracleService.FunctionOracleListen(errChan3)
	functionConsumerService := NewFunctionConsumerEventService(client, l.db, eth.MUMBAI_TESTNET)
	functionConsumerService.FunctionConsumerListen(errChan4)
	for {
		select {
		case err1 := <-errChan:
			logger.Errorf("speol watch event error, error1 is : %s", err1)
			return err1
		case err2 := <-errChan2:
			logger.Errorf("speol watch event error, error2 is : %s", err2)
			return err2
		case err3 := <-errChan3:
			logger.Errorf("speol watch event error, error3 is : %s", err3)
			return err3
		case err4 := <-errChan4:
			logger.Errorf("speol watch event error, error4 is : %s", err4)
			return err4
		}
	}
	return nil
}
