package initialization

import (
	"github.com/joho/godotenv"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/handler"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/utils"
	"os"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	utils.InitLogger()
	InitDB()
	httpHandler := handler.NewHandlerServer()
	userService := aline.NewUserService()
	application.SetBean[*aline.UserService]("userService", userService)

	err = handler.NewHttpService(*httpHandler, os.Getenv("PORT")).StartHttpServer()
	if err != nil {
		panic(err)
	}
}
