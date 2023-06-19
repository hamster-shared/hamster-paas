package node

import (
	"context"
	"database/sql"
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
		err := ol.db.Model(order.Order{}).Where("status = ?", order.PaymentPending).Order("order_time asc").Find(&orderList).Error
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
			err := ol.db.Model(order.ReceiptRecords{}).Where("amount = ? and pay_address = ? and receive_address = ? and pay_time > ? and pay_time < ?", orderInfo.Amount, orderInfo.PayAddress, orderInfo.ReceiveAddress, orderInfo.OrderTime.Time, orderInfo.OrderTime.Time.Add(time.Hour)).Order("pay_time asc").Find(&receiptRecords).Error
			if err != nil {
				logger.Errorf("Failed to query the ReceiptRecords: %s", err)
				return
			}
			begin := ol.db.Begin()
			if len(receiptRecords) >= 1 {
				orderInfo.Status = order.Paid
				orderInfo.PayTx = receiptRecords[0].PayTx
				//
				var orderNode order.OrderNode
				err := begin.Model(order.OrderNode{}).Where("order_id = ? and user_id = ?", orderInfo.OrderId, orderInfo.UserId).Find(&orderNode).Error
				if err != nil {
					logger.Errorf("Failed to query OrderNode: %s", err)
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
					PaymentPerMonth:   decimal.Decimal{},
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
	EntryID, err := c.AddFunc("*/15 * * * * *", func() {
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
		//fmt.Println("当前块高度:", currentBlockHeight)
		//block, err := ol.client.BlockByNumber(context.Background(), currentBlockHeight)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//for _, tx := range block.Transactions() {
		//	fmt.Printf("block.Transactions 交易哈希：%s\n", tx.Hash().Hex())
		//	//fmt.Printf("block.Transactions 发送方地址：%s\n", tx.from.Load())
		//	fmt.Printf("block.Transactions 接收方地址：%s\n", tx.To().Hex())
		//	fmt.Printf("block.Transactions 交易金额：%s\n", tx.Value().String())
		//}

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

		var records []order.ReceiptRecords
		// 处理事件日志
		for _, log := range logs {
			var receiptRecords order.ReceiptRecords
			// 获取交易时间戳
			block, err := ol.client.BlockByHash(context.Background(), log.BlockHash)
			if err == nil {
				timestamp := time.Unix(int64(block.Time()), 0)
				receiptRecords.PayTime = timestamp
				fmt.Printf("交易时间：%s\n", timestamp.String())
			}

			from := common.BytesToAddress(log.Topics[1].Bytes())
			to := common.BytesToAddress(log.Topics[2].Bytes())
			amount := new(big.Int).SetBytes(log.Data)
			amountDecimal := decimal.NewFromBigInt(amount, 0).Div(decimal.NewFromInt(1e6))
			receiptRecords.PayTx = log.TxHash.Hex()
			fmt.Printf("交易哈希：%s\n", log.TxHash.Hex())
			receiptRecords.PayAddress = from.Hex()
			fmt.Printf("发送方地址：%s\n", from.Hex())
			receiptRecords.ReceiveAddress = to.Hex()
			fmt.Printf("接收方地址：%s\n", to.Hex())
			receiptRecords.Amount = amountDecimal
			fmt.Printf("交易金额：%s\n", amountDecimal.String())
			if strings.Contains(addressString, receiptRecords.ReceiveAddress) {
				records = append(records, receiptRecords)
			}
		}
		blackHeight.BlackHeight = int64(currentBlockHeight + 1)
		if len(records) < 1 {
			err = ol.db.Model(&blackHeight).Updates(&blackHeight).Error
			return
		}
		begin := ol.db.Begin()
		err = begin.Model(order.ReceiptRecords{}).Create(&records).Error
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
		}
		begin.Commit()
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
		PingTimeout:  time.Second * 5,
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

	server.OnEvent("/", "order_status", func(s socketIo.Conn, orderId int) {
		logger.Infof("orderId: %d\n", orderId)
		var orderData order.Order
		for {
			err := ol.db.Model(order.Order{}).Where("id = ?", orderId).First(&orderData).Error
			if err != nil {
				time.Sleep(time.Second * 5)
				continue
			}
			if orderData.Status == order.Cancelled || orderData.Status == order.Paid {
				break
			}
			if orderData.Status == order.PaymentPending {
				time.Sleep(time.Second * 5)
			}
		}
		s.Emit("order_result", orderData.Status)
	})

	server.OnError("/", func(s socketIo.Conn, err error) {
		logger.Errorf("socket meet error: %v\n", err)
	})

	server.OnDisconnect("/", func(s socketIo.Conn, reason string) {
		logger.Errorf("socket closed, reason is: %s", reason)
	})
	return server
}
