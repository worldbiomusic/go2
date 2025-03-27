package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ReadEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env 파일 읽기 실패")
	}

	value := os.Getenv(key)
	return value
}
