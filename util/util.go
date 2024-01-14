package util

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv(path string) error {
	// .env 파일 로드
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

func NewToken(path string) string {
	LoadEnv(path)
	result := os.Getenv("TOKEN")
	return result
}
