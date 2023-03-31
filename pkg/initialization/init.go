package initialization

import (
	"fmt"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/handler"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/service"
	"hamster-paas/pkg/service/nginx_log_parse"
	"hamster-paas/pkg/utils/logger"
	"os"

	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("error loading .env file: %s", err))
	}
	logger.InitLogger()
	InitDB()
	aline.NewAlineRpc().Init()
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		panic("application get db failed")
	}
	chainLinkRequestService := service.NewChainLinkRequestService(db)
	application.SetBean[*service.ChainLinkRequestService]("chainLinkRequestService", chainLinkRequestService)
	chainLinkSubscriptionService := service.NewChainLinkSubscriptionService(db)
	application.SetBean[*service.ChainLinkSubscriptionService]("chainLinkSubscriptionService", chainLinkSubscriptionService)
	chainLinkConsumerService := service.NewChainLinkConsumerService(db)
	application.SetBean[*service.ChainLinkConsumerService]("chainLinkConsumerService", chainLinkConsumerService)
	chainLinkDepositService := service.NewChainLinkDepositService(db)
	application.SetBean[*service.ChainLinkDepositService]("chainLinkDepositService", chainLinkDepositService)
	httpHandler := handler.NewHandlerServer()
	err = handler.NewHttpService(*httpHandler, os.Getenv("PORT")).StartHttpServer()
	if err != nil {
		panic(err)
	}
	nginx_log_parse.InitMeiliSearch()
}
