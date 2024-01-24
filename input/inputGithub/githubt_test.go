package inputGithub

import (
	"bytes"
	"fineC/util"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewGithubClient(t *testing.T) {
	// client 객체 생성을 확인하기 위한 mock 객체 생성
	token := util.NewToken("../'.env")
	client := NewGithubClient(token, "kubernetes/kubernetes")
	test := assert.New(t)

	// 테스트 로직
	test.Equal(token, client.Token)
	test.Equal("https://api.github.com/repos/kubernetes/kubernetes/issues", client.Url)
}

func TestCallAPI(t *testing.T) {
	// 테스트 서버 설정
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`test response`))
	}))
	defer ts.Close()

	// GitHubClient 인스턴스 생성
	client := &GitHubClient{
		Url: ts.URL,
	}

	resp, _ := client.CallAPI(1)

	if !bytes.Equal(resp, []byte(`test response`)) {
		t.Errorf("Expected 'test response', got '%s'", resp)
	}
}

func TestGitHubClient_Crawler(t *testing.T) {
	// Github client 객체를 생성하기 위한 mock 객체 생성
	// 테스트 서버 설정
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`test response`))
	}))
	defer ts.Close()
	mock := &GitHubClient{
		Url: ts.URL,
	}
	ch := make(chan []byte, 2)
	close(ch)
	mock.Crawling(ch)

	test := assert.New(t)

	for data := range ch {
		test.Equal(data, []byte(`test response`))
	}

	//test.
}
