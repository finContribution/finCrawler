package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
)

func main() {
	// GitHub 개인 액세스 토큰 설정
	token := os.Getenv("TOKEN")

	// OAuth2 토큰 생성
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(context.Background(), ts)

	// GitHub 클라이언트 생성
	client := github.NewClient(tc)

	// GitHub 사용자 정보 조회 예제
	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		fmt.Printf("Error getting user: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("GitHub 사용자 정보:\n")
	fmt.Printf("이름: %s\n", *user.Name)
	fmt.Printf("로그인: %s\n", *user.Login)
	fmt.Printf("블로그: %s\n", *user.Blog)

}
