package util

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() error {
	// .env 파일 로드
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func NewToken() string {
	LoadEnv()
	result := os.Getenv("TOKEN")
	return result
}
