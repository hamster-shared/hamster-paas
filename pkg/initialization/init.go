package initialization

import (
	"github.com/joho/godotenv"
	"hamster-paas/pkg/aline/service"
	"hamster-paas/pkg/api"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/handler"
	"hamster-paas/pkg/logger"
	"os"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	logger.InitLogger()
	InitDB()
	httpHandler := handler.NewHandlerServer()
	userService := service.NewUserService()
	application.SetBean[*service.UserService]("userService", userService)

	api.NewHttpService(*httpHandler, os.Getenv("PORT")).StartHttpServer()
}
