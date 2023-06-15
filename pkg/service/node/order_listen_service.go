package node

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"hamster-paas/pkg/models/order"
	"hamster-paas/pkg/utils/logger"
	"log"
	"os"
	"time"
)

type OrderListeningService struct {
	db                   *gorm.DB
	erc20ContractAddress common.Address
	client               *ethclient.Client
}

func NewOrderListeningService(db *gorm.DB) *OrderListeningService {
	return &OrderListeningService{
		db: db,
	}
}

func (ol *OrderListeningService) StartOrderListening() {
	//cron.WithSeconds()	秒级操作
	//cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger))	函数没执行完就跳过本次函数
	//cron.WithLogger(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags)))	打印任务日志
	c := cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)), cron.WithLogger(
		cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
	EntryID, err := c.AddFunc("*/5 * * * * *", func() {
		fmt.Println(time.Now(), "支付中订单扫描")
		var orderList []order.Order
		err := ol.db.Model(order.Order{}).Where("status = ?", order.PaymentPending).Order("order_time asc").Find(&orderList).Error
		if err != nil {
			logger.Errorf("Failed to query the order in payment: %s", err)
			return
		}
		if len(orderList) < 1 {
			logger.Infof("No orders in payment", err)
			return
		}
		//获取地址中所有的Transfer Event

		for _, order := range orderList {
			//查询获取订单中地址
			//
			fmt.Println(order)
		}

	})
	if err != nil {
		logger.Errorf("order cron start failed, EntryID: %s, err: %s", EntryID, err)
	}
	c.Start()
}
