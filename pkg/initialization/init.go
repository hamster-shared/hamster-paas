package initialization

import (
	"fmt"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/handler"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/service"
	"hamster-paas/pkg/service/nginx_log_parse"
	service2 "hamster-paas/pkg/service/node"
	"hamster-paas/pkg/utils/logger"
	"os"

	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		// panic(fmt.Errorf("error loading .env file: %s", err))
		// 如果获取不到的话，也没事，可能是从 docker 或 k8s 里面启动的
		fmt.Println("warning: dont load .env file")
	}
	fmt.Println("init logger")
	logger.InitLogger()
	fmt.Println("init db")
	InitDB()
	fmt.Println("init aline db")
	aline.NewAlineRpc().Init()
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		panic("application get db failed")
	}
	fmt.Println("chainlink request service")
	chainLinkRequestService := service.NewChainLinkRequestService(db)
	application.SetBean[*service.ChainLinkRequestService]("chainLinkRequestService", chainLinkRequestService)

	fmt.Println("chainlink subscription service")
	chainLinkSubscriptionService := service.NewChainLinkSubscriptionService(db)
	application.SetBean[*service.ChainLinkSubscriptionService]("chainLinkSubscriptionService", chainLinkSubscriptionService)

	fmt.Println("chainlink consumer service")
	chainLinkConsumerService := service.NewChainLinkConsumerService(db)
	application.SetBean[*service.ChainLinkConsumerService]("chainLinkConsumerService", chainLinkConsumerService)

	fmt.Println("chainlink deposit service")
	chainLinkDepositService := service.NewChainLinkDepositService(db)
	application.SetBean[*service.ChainLinkDepositService]("chainLinkDepositService", chainLinkDepositService)
	application.SetBean("chainLinkDashboardService", service.NewChainLinkDashboardService(db))

	fmt.Println("chainlink pool service")
	chainLinkPoolService := service.NewPoolService()
	application.SetBean[*service.PoolService]("chainLinkPoolService", chainLinkPoolService)

	fmt.Println("chainlink long link pool service")
	longLinkPoolService := service.NewLongLinkPoolService()
	application.SetBean[*service.LongLinkPoolService]("longLinkPoolService", longLinkPoolService)

	fmt.Println("rpc service")
	application.SetBean("rpcService", service.NewRpcService(db))

	fmt.Println("middle ware service")
	application.SetBean("middleWareService", service.NewMiddleWareService(db))

	fmt.Println("meili search service")
	nginx_log_parse.InitMeiliSearch()

	fmt.Println("oracle listener service")
	oracleListener := service.NewOracleListener(db)
	oracleListener.StartListen()
	application.SetBean("oracleListener", oracleListener)

	application.SetBean("nodeService", service2.NewNodeService(db))
	application.SetBean("orderService", service2.NewOrderService(db))
	application.SetBean("resourceStandardService", service2.NewResourceStandardService(db))

	fmt.Println("handler server")
	httpHandler := handler.NewHandlerServer()
	err = handler.NewHttpService(*httpHandler, os.Getenv("PORT")).StartHttpServer()
	if err != nil {
		panic(err)
	}
}
