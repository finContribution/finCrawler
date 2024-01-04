package main

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	// .env 파일 로드
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}

func init() {
	err := LoadEnv()

	if err != nil {
		fmt.Printf("Sommting wrong, reson %s", err)
		panic(err)
	}

}
