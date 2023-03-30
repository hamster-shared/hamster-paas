package initialization

import (
	"fmt"
	"gorm.io/gorm"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/handler"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/service"
	"hamster-paas/pkg/utils/logger"
	"os"

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
	chainLinkSubscriptionService := service.NewChainLinkSubscriptionService(db)
	application.SetBean[*service.ChainLinkRequestService]("chainLinkRequestService", chainLinkRequestService)
	application.SetBean[*service.ChainLinkSubscriptionService]("chainLinkSubscriptionService", chainLinkSubscriptionService)
	httpHandler := handler.NewHandlerServer()
	err = handler.NewHttpService(*httpHandler, os.Getenv("PORT")).StartHttpServer()
	if err != nil {
		panic(err)
	}
}
