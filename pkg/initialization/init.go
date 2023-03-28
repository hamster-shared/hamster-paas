package initialization

import "github.com/joho/godotenv"

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	InitDB()
}
