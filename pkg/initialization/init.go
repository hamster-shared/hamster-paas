package initialization

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/handler"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/service"
	"hamster-paas/pkg/utils/logger"
	"os"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
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

	httpHandler := handler.NewHandlerServer()
	err = handler.NewHttpService(*httpHandler, os.Getenv("PORT")).StartHttpServer()
	if err != nil {
		panic(err)
	}
}
