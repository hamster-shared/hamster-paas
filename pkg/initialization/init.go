package initialization

import (
	"fmt"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/handler"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/service"
	"hamster-paas/pkg/service/nginx_log_parse"
	service2 "hamster-paas/pkg/service/node"
	"hamster-paas/pkg/service/zan"
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
	zanApiEndpoint := os.Getenv("ZAN_API_ENDPOINT")
	zanDefaultAccessToken := os.Getenv("ZAN_DEFAULT_ACCESS_TOKEN")
	zanClientId := os.Getenv("ZAN_CLIENT_ID")
	zanPrivateKeyPATH := os.Getenv("ZAN_PRIVATE_KEY_PATH")
	//zanClient := zan.NewZanClient("http://webtcapi.unchartedw3s.com", "478f53734d284889a6a0fbfc648cc061", "2def8d1826884fdd896508d078b26a91", "/Users/mohaijiang/IdeaProjects/blockchain/hamster-paas/rsa_private_key_pkcs8.pem")
	zanClient := zan.NewZanClient(zanApiEndpoint, zanDefaultAccessToken, zanClientId, zanPrivateKeyPATH)

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
	application.SetBean("rpcService", service.NewRpcService(db, zanClient))

	fmt.Println("middle ware service")
	application.SetBean("middleWareService", service.NewMiddleWareService(db, zanClient))

	fmt.Println("meili search service")
	nginx_log_parse.InitMeiliSearch()

	fmt.Println("oracle listener service")
	oracleListener := service.NewOracleListener(db)
	oracleListener.StartListen()
	application.SetBean("oracleListener", oracleListener)

	application.SetBean("nodeService", service2.NewNodeService(db))
	application.SetBean("orderService", service2.NewOrderService(db))
	application.SetBean("resourceStandardService", service2.NewResourceStandardService(db))
	listeningService := service2.NewOrderListeningService(os.Getenv("TOKEN_ADDRESS"), db)
	listeningService.StartOrderListening()
	listeningService.StartScanBlockInformation()

	zanService := service.NewZanService(zanClient, db)
	application.SetBean("zanService", zanService)
	fmt.Println("handler server")
	httpHandler := handler.NewHandlerServer()
	err = handler.NewHttpService(*httpHandler, os.Getenv("PORT"), listeningService.GetOrderWebSocket()).StartHttpServer()
	if err != nil {
		panic(err)
	}
}
