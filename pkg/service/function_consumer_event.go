package service

import (
	"context"
	"encoding/hex"
	"fmt"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/rpc/eth"
	"hamster-paas/pkg/service/contract"
	"hamster-paas/pkg/utils/logger"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type FunctionConsumerEventService struct {
	client  *ethclient.Client
	db      *gorm.DB
	network eth.EthNetwork
}

func NewFunctionConsumerEventService(client *ethclient.Client, db *gorm.DB, network eth.EthNetwork) *FunctionConsumerEventService {
	return &FunctionConsumerEventService{
		client:  client,
		db:      db,
		network: network,
	}
}

func (b *FunctionConsumerEventService) FunctionConsumerListen(errorChan chan error) {
	go func() {
		err := b.oCRResponseListen()
		errorChan <- err
	}()
}

func (b *FunctionConsumerEventService) oCRResponseListen() error {
	oracleContractAddress := common.HexToAddress("")
	// 定义查询过滤器
	query := ethereum.FilterQuery{
		Topics: [][]common.Hash{
			{
				crypto.Keccak256Hash([]byte("OCRResponse(bytes32,bytes,bytes)")),
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
	contractFilter, err := contract.NewFunctionConsumer(oracleContractAddress, b.client)
	if err != nil {
		logger.Errorf("create Oracle Request failed: %s", err)
		return err
	}
	for {
		select {
		case err := <-sub.Err():
			logger.Errorf("subscribe Oracle Request event failed: %s", err)
			return err
		case vLog := <-logs:
			logger.Info("start watch Oracle Request event send email")
			data, err := contractFilter.ParseOCRResponse(vLog)
			if err == nil {
				logger.Debugf("已经监听到该事件")
				handlerData(data, b.db)
			}
		}
	}
}

func handlerData(data *contract.FunctionConsumerOCRResponse, db *gorm.DB) {
	var eventData models.FunctionConsumerEvent
	requestIdData := fmt.Sprintf("0x%s", hex.EncodeToString(data.RequestId[:]))
	logger.Debugf("++++++++++++++++++++")
	fmt.Printf("request id is:%s", requestIdData)
	logger.Debugf("++++++++++++++++++++")
	var result string
	numData, err := strconv.ParseInt(hexToString(data.Result), 16, 64)
	if err != nil {
		result = string(data.Result)
	} else {
		result = strconv.Itoa(int(numData))
	}
	logger.Debugf("*****************")
	logger.Debugf(result)
	logger.Debugf("*****************")
	logger.Debugf("#################")
	logger.Debugf(string(data.Err))
	logger.Debugf("#################")
	eventData.Created = time.Now()
	eventData.RequestId = requestIdData
	eventData.Result = result
	eventData.ErrorInfo = string(data.Err)
	db.Create(&eventData)
}

func hexToString(hex []byte) string {
	s := ""
	for _, b := range hex {
		s += fmt.Sprintf("%02x", b)
	}
	return s
}
