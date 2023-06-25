package node

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	socketIo "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/robfig/cron/v3"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"hamster-paas/pkg/models/node"
	"hamster-paas/pkg/models/order"
	"hamster-paas/pkg/utils"
	"hamster-paas/pkg/utils/logger"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"
)

type OrderListeningService struct {
	db                   *gorm.DB
	erc20ContractAddress common.Address
	client               *ethclient.Client
}

func NewOrderListeningService(erc20ContractAddress string, db *gorm.DB) *OrderListeningService {
	client, err := ethclient.Dial(os.Getenv("NODE_URL"))
	if err != nil {
		logger.Errorf("connect NODE_URL failed: %s", err)
		panic("application get NODE_URL client failed")
	}
	return &OrderListeningService{
		db:                   db,
		erc20ContractAddress: common.HexToAddress(erc20ContractAddress),
		client:               client,
	}
}

func (ol *OrderListeningService) StartOrderListening() {
	//cron.WithSeconds()	秒级操作
	//cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger))	函数没执行完就跳过本次函数
	//cron.WithLogger(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags)))	打印任务日志
	c := cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)), cron.WithLogger(
		cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
	EntryID, err := c.AddFunc("*/5 * * * * *", func() {
		var orderList []order.Order
		err := ol.db.Model(order.Order{}).Where("status = ?", order.PaymentPending).Order("order_time desc").Find(&orderList).Error
		if err != nil {
			logger.Errorf("Failed to query the order in payment: %s", err)
			return
		}
		if len(orderList) < 1 {
			logger.Info("There are no orders in payment")
			return
		}
		for _, orderInfo := range orderList {
			//查询获取订单中地址
			var receiptRecords []order.ReceiptRecords
			err := ol.db.Model(order.ReceiptRecords{}).Where("amount = ? and receive_address = ? and pay_time > ? and pay_time < ? and order_id is null", orderInfo.Amount, orderInfo.ReceiveAddress, orderInfo.OrderTime.Time, orderInfo.OrderTime.Time.Add(time.Hour)).Order("pay_time desc").Find(&receiptRecords).Error
			if err != nil {
				logger.Errorf("Failed to query the ReceiptRecords: %s", err)
				return
			}
			logger.Infof("The number of transactions waiting to be bound is %d \n", len(receiptRecords))
			begin := ol.db.Begin()
			data, err := AccountBalance(orderInfo.ReceiveAddress)
			if err != nil {
				logger.Errorf("get balance failed: %s", err)
				continue
			}
			balance, err := decimal.NewFromString(data)
			if err != nil {
				logger.Errorf("balance to decimal failed: %s", err)
				continue
			}
			addressInitBalance, err := decimal.NewFromString(orderInfo.AddressInitBalance.String)
			if err != nil {
				logger.Errorf("addressInitBalance to decimal failed: %s", err)
				continue
			}
			amount, err := decimal.NewFromString(orderInfo.Amount.String)
			if err != nil {
				logger.Errorf("amount to decimal failed: %s", err)
				continue
			}
			cmp := balance.Cmp(addressInitBalance.Add(amount))
			if len(receiptRecords) >= 1 && cmp == 0 {
				var orderDb order.Order
				err := ol.db.Model(&order.Order{}).Where("pay_tx = ?", receiptRecords[0].PayTx).First(&orderDb).Error
				if !errors.Is(gorm.ErrRecordNotFound, err) {
					if orderInfo.OrderTime.Time.Add(time.Hour).Before(time.Now()) {
						orderInfo.Status = order.Cancelled
					}
				} else {
					orderInfo.Status = order.Paid
					orderInfo.PayTx = receiptRecords[0].PayTx
					orderInfo.PayAddress = receiptRecords[0].PayAddress

					var orderNode order.OrderNode
					err = begin.Model(order.OrderNode{}).Where("order_id = ? and user_id = ?", orderInfo.Id, orderInfo.UserId).Find(&orderNode).Error
					if err != nil {
						logger.Errorf("Failed to query OrderNode: %s", err)
						begin.Callback()
						return
					}
					err := begin.Model(order.ReceiptRecords{}).Where("id = ?", receiptRecords[0].Id).Update("order_id", orderInfo.Id).Error
					if err != nil {
						logger.Errorf("Failed to update ReceiptRecords: %s", err)
						begin.Callback()
						return
					}
					RPCNode := node.RPCNode{
						Name:          orderNode.NodeName,
						UserId:        orderNode.UserId,
						ChainProtocol: node.ChainProtocol(orderNode.Protocol),
						Status:        node.Initializing,
						PublicIp:      "",
						Region:        orderNode.Region,
						LaunchTime:    orderInfo.OrderTime,
						Resource:      orderNode.Resource,
						ChainVersion:  "",
						NextPaymentDate: sql.NullTime{
							Time:  orderInfo.OrderTime.Time.AddDate(0, 1, 0),
							Valid: true,
						},
						PaymentPerMonth: sql.NullString{
							String: "0.00",
							Valid:  true,
						},
						RemainingSyncTime: "",
						CurrentHeight:     0,
						BlockTime:         "",
						HttpEndpoint:      "",
						WebsocketEndpoint: "",
						Created: sql.NullTime{
							Time:  time.Now(),
							Valid: true,
						},
					}
					err = begin.Model(node.RPCNode{}).Create(&RPCNode).Error
					if err != nil {
						logger.Errorf("Failed to Create OrderNode: %s", err)
						begin.Callback()
						return
					} else {
						utils.SendEmailForNodeCreate(RPCNode)
					}
				}
			} else {
				if orderInfo.OrderTime.Time.Add(time.Hour).Before(time.Now()) {
					orderInfo.Status = order.Cancelled
				}
			}
			err = begin.Model(&orderInfo).Updates(&orderInfo).Error
			if err != nil {
				logger.Errorf("Failed to Updates Order: %s", err)
				begin.Callback()
				return
			} else {
				begin.Commit()
			}
		}
	})
	if err != nil {
		logger.Errorf("StartOrderListening start failed, EntryID: %s, err: %s", EntryID, err)
	}
	c.Start()
}

func (ol *OrderListeningService) StartScanBlockInformation() {
	c := cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)), cron.WithLogger(
		cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
	EntryID, err := c.AddFunc("*/5 * * * * *", func() {
		var blackHeight order.BlackHeight
		err := ol.db.Model(order.BlackHeight{}).Where("event_type = ?", "Transfer").First(&blackHeight).Error
		if err != nil {
			logger.Errorf("Failed to query the BlackHeight in db: %s", err)
			return
		}

		// 获取当前块高度
		currentBlockHeight, err := ol.client.BlockNumber(context.Background())
		if err != nil {
			logger.Errorf("Failed to query the BlockNumber in client: %s", err)
			return
		}
		if int64(currentBlockHeight) <= blackHeight.BlackHeight {
			return
		}

		//扫描事件块
		query := ethereum.FilterQuery{
			Addresses: []common.Address{ol.erc20ContractAddress},
			Topics: [][]common.Hash{
				{
					crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")),
				},
			},
			FromBlock: big.NewInt(blackHeight.BlackHeight),
			ToBlock:   big.NewInt(int64(currentBlockHeight)),
		}
		logger.Infof("Scan transaction events from %d to %d\n", blackHeight.BlackHeight, currentBlockHeight)
		logs, err := ol.client.FilterLogs(context.Background(), query)
		if err != nil {
			logger.Errorf("Failed to FilterLogs: %s", err)
			return
		}
		var addresses []string
		err = ol.db.Model(node.UserChargeAccount{}).Pluck("address", &addresses).Error
		if err != nil {
			logger.Errorf("Failed to query address in db: %s", err)
			return
		}
		addressString := strings.Join(addresses, "\n")

		records := []map[string]interface{}{}
		// 处理事件日志
		for _, log := range logs {
			var receiptRecords order.ReceiptRecords
			// 获取交易时间戳
			block, err := ol.client.BlockByHash(context.Background(), log.BlockHash)
			if err == nil {
				timestamp := time.Unix(int64(block.Time()), 0)
				receiptRecords.PayTime = timestamp
				receiptRecords.PayTimeUTC = timestamp.Add(-8 * time.Hour)
				fmt.Printf("交易时间：%s\n", timestamp.String())
			}
			headerByHash, err := ol.client.HeaderByHash(context.Background(), log.BlockHash)
			if err == nil {
				receiptRecords.BlackHeight = headerByHash.Number.Int64()
				fmt.Println("Block Number:", headerByHash.Number.Int64())
			}

			from := common.BytesToAddress(log.Topics[1].Bytes())
			to := common.BytesToAddress(log.Topics[2].Bytes())
			amount := new(big.Int).SetBytes(log.Data)
			amountDecimal := decimal.NewFromBigInt(amount, 0).Div(decimal.NewFromInt(1e6))
			receiptRecords.PayTx = log.TxHash.Hex()
			fmt.Printf("交易哈希：%s\n", log.TxHash.Hex())
			var dbReceiptRecords order.ReceiptRecords
			err = ol.db.Model(&order.ReceiptRecords{}).Where("pay_tx = ?", receiptRecords.PayTx).First(&dbReceiptRecords).Error
			if !errors.Is(gorm.ErrRecordNotFound, err) {
				logger.Errorf("pay_tx %s already exists in the db! err is %s\n", receiptRecords.PayTx, err)
				continue
			}
			receiptRecords.PayAddress = from.Hex()
			fmt.Printf("发送方地址：%s\n", from.Hex())
			receiptRecords.ReceiveAddress = to.Hex()
			fmt.Printf("接收方地址：%s\n", to.Hex())
			receiptRecords.Amount = amountDecimal
			fmt.Printf("交易金额：%s\n", amountDecimal.String())
			if strings.Contains(addressString, receiptRecords.ReceiveAddress) {
				data := map[string]interface{}{
					"black_height":    receiptRecords.BlackHeight,
					"pay_address":     receiptRecords.PayAddress,
					"receive_address": receiptRecords.ReceiveAddress,
					"amount":          receiptRecords.Amount,
					"pay_tx":          receiptRecords.PayTx,
					"pay_time":        receiptRecords.PayTime,
					"pay_time_utc":    receiptRecords.PayTimeUTC,
				}
				records = append(records, data)
			}
		}
		blackHeight.BlackHeight = int64(currentBlockHeight + 1)
		if len(records) < 1 {
			err = ol.db.Model(&blackHeight).Updates(&blackHeight).Error
			return
		}
		begin := ol.db.Begin()
		err = begin.Model(&order.ReceiptRecords{}).Create(&records).Error
		if err != nil {
			logger.Errorf("Failed to add ReceiptRecords to db: %s", err)
			begin.Callback()
			return
		}

		err = begin.Model(&blackHeight).Updates(&blackHeight).Error
		if err != nil {
			logger.Errorf("Failed to Updates blackHeight to db: %s", err)
			begin.Callback()
			return
		} else {
			begin.Commit()
		}
	})
	if err != nil {
		logger.Errorf("StartScanBlockInformation start failed, EntryID: %s, err: %s", EntryID, err)
	}
	c.Start()
}

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func (ol *OrderListeningService) GetOrderWebSocket() *socketIo.Server {
	server := socketIo.NewServer(&engineio.Options{
		PingTimeout:  time.Hour,
		PingInterval: time.Second * 2,
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})

	server.OnConnect("/", func(s socketIo.Conn) error {
		s.SetContext("")
		logger.Infof("connected session id: %s", s.ID())
		return nil
	})

	server.OnEvent("/page", "order_status_page", func(s socketIo.Conn, orderId int) {
		logger.Infof("order_status_page orderId: %d\n", orderId)
		status := ol.PollingGetOrderStatus(orderId)
		logger.Infof("order_status_page orderId: %d Sending results to the client---> %d\n", orderId, status)
		s.Emit("order_result", status)
	})

	server.OnEvent("/", "order_status_model", func(s socketIo.Conn, orderId int) {
		logger.Infof("order_status_model orderId: %d\n", orderId)
		status := ol.PollingGetOrderStatus(orderId)
		logger.Infof("order_status_model orderId: %d Sending results to the client---> %d\n", orderId, status)
		s.Emit("order_result", status)
	})

	server.OnError("/", func(s socketIo.Conn, err error) {
		logger.Errorf("socket meet error: %v\n", err)
	})

	server.OnDisconnect("/", func(s socketIo.Conn, reason string) {
		logger.Errorf("socket closed, reason is: %s", reason)
	})
	return server
}

func (ol *OrderListeningService) PollingGetOrderStatus(orderId int) int {
	var orderData order.Order
	for {
		err := ol.db.Model(order.Order{}).Where("id = ?", orderId).First(&orderData).Error
		if err != nil {
			logger.Errorf("orderId: %d query fail, err is : %s\n", orderId, err)
			time.Sleep(time.Second * 3)
			continue
		}
		if orderData.Status == order.Cancelled || orderData.Status == order.Paid {
			logger.Infof("orderId: %d query end\n", orderId)
			break
		}
		if orderData.Status == order.PaymentPending {
			logger.Infof("orderId: %d query continue\n", orderId)
			time.Sleep(time.Second * 3)
		}
	}
	return int(orderData.Status)
}
