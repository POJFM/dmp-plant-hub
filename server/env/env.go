package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Process(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env, make sure there is one")
	}
	return os.Getenv(key)
}
