package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Process(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env, make sure there is one")
	}
	return os.Getenv(key)
}
