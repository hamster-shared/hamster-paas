package initialization

import (
	"hamster-paas/pkg/logger"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	logger.InitLogger()
	InitDB()
}
