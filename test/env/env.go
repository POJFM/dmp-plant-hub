package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Process(key string) string {
	err := godotenv.Load("env/.env")
	if err != nil {
		log.Fatal("Error loading .env")
	}
	return os.Getenv(key)
}
