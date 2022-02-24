package env

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func Process(key string) string {
	err := godotenv.Load(filepath.Join("./env/", ".env"), filepath.Join(".env"))
	if err != nil {
		log.Fatal("Error loading .env, make shure there is one")
	}
	return os.Getenv(key)
}
