package initialization

import (
	"fmt"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/handler"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/service"
	service2 "hamster-paas/pkg/service/node"
	"hamster-paas/pkg/service/zan"
	"hamster-paas/pkg/utils/logger"
	"os"

	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

// 初始化函数
func Init() {
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		// panic(fmt.Errorf("error loading .env file: %s", err))
		// 如果获取不到的话，也没事，可能是从 docker 或 k8s 里面启动的
		fmt.Println("warning: dont load .env file")
	}
	// 初始化日志
	fmt.Println("init logger")
	logger.InitLogger()
	// 初始化数据库
	fmt.Println("init db")
	InitDB()
	// 初始化 aline 数据库
	fmt.Println("init aline db")
	aline.NewAlineRpc().Init()
	// 获取数据库
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		panic("application get db failed")
	}
	// 获取 zan 环境变量
	zanApiEndpoint := os.Getenv("ZAN_API_ENDPOINT")
	zanDefaultAccessToken := os.Getenv("ZAN_DEFAULT_ACCESS_TOKEN")
	zanClientId := os.Getenv("ZAN_CLIENT_ID")
	zanPrivateKeyPATH := os.Getenv("ZAN_PRIVATE_KEY_PATH")
	// zanClient := zan.NewZanClient("http://webtcapi.unchartedw3s.com", "478f53734d284889a6a0fbfc648cc061", "2def8d1826884fdd896508d078b26a91", "/Users/mohaijiang/IdeaProjects/blockchain/hamster-paas/rsa_private_key_pkcs8.pem")
	zanClient := zan.NewZanClient(zanApiEndpoint, zanDefaultAccessToken, zanClientId, zanPrivateKeyPATH)

	// 初始化 chainlink request service
	fmt.Println("chainlink request service")
	chainLinkRequestService := service.NewChainLinkRequestService(db)
	application.SetBean[*service.ChainLinkRequestService]("chainLinkRequestService", chainLinkRequestService)

	// 初始化 chainlink subscription service
	fmt.Println("chainlink subscription service")
	chainLinkSubscriptionService := service.NewChainLinkSubscriptionService(db)
	application.SetBean[*service.ChainLinkSubscriptionService]("chainLinkSubscriptionService", chainLinkSubscriptionService)

	// 初始化 chainlink consumer service
	fmt.Println("chainlink consumer service")
	chainLinkConsumerService := service.NewChainLinkConsumerService(db)
	application.SetBean[*service.ChainLinkConsumerService]("chainLinkConsumerService", chainLinkConsumerService)

	// 初始化 chainlink deposit service
	fmt.Println("chainlink deposit service")
	chainLinkDepositService := service.NewChainLinkDepositService(db)
	application.SetBean[*service.ChainLinkDepositService]("chainLinkDepositService", chainLinkDepositService)
	application.SetBean("chainLinkDashboardService", service.NewChainLinkDashboardService(db))

	// 初始化 chainlink pool service
	fmt.Println("chainlink pool service")
	chainLinkPoolService := service.NewPoolService()
	application.SetBean[*service.PoolService]("chainLinkPoolService", chainLinkPoolService)

	// 初始化 chainlink long link pool service
	fmt.Println("chainlink long link pool service")
	longLinkPoolService := service.NewLongLinkPoolService()
	application.SetBean[*service.LongLinkPoolService]("longLinkPoolService", longLinkPoolService)

	// 初始化 rpc service
	fmt.Println("rpc service")
	application.SetBean("rpcService", service.NewRpcService(db, zanClient))

	// 初始化 middle ware service
	fmt.Println("middle ware service")
	application.SetBean("middleWareService", service.NewMiddleWareService(db, zanClient))

	// 初始化 meili search service
	fmt.Println("meili search service")
	// nginx_log_parse.InitMeiliSearch()

	// 初始化 oracle listener service
	fmt.Println("oracle listener service")
	oracleListener := service.NewOracleListener(db)
	oracleListener.StartListen()
	application.SetBean("oracleListener", oracleListener)

	// 初始化 node service
	application.SetBean("nodeService", service2.NewNodeService(db))
	// 初始化 order service
	application.SetBean("orderService", service2.NewOrderService(db))
	// 初始化 resource standard service
	application.SetBean("resourceStandardService", service2.NewResourceStandardService(db))
	// 初始化 order listening service
	listeningService := service2.NewOrderListeningService(os.Getenv("TOKEN_ADDRESS"), db)
	listeningService.StartOrderListening()
	listeningService.StartScanBlockInformation()

	// 初始化 zan service
	zanService := service.NewZanService(zanClient, db)
	application.SetBean("zanService", zanService)

	icpService := service.NewIcpService()
	application.SetBean("icpService", icpService)
	// 初始化 handler server
	fmt.Println("handler server")
	httpHandler := handler.NewHandlerServer()
	err = handler.NewHttpService(*httpHandler, os.Getenv("PORT"), listeningService.GetOrderWebSocket()).StartHttpServer()
	if err != nil {
		panic(err)
	}
}
