package input

import (
	"fineC/util"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
.env 파일에 추가한 token 정보를 가져옵니다 token 생성관련 내용은 링크 참조해주세요
https://docs.github.com/en/rest/authentication/authenticating-to-the-rest-api?apiversion=2022-11-28&apiVersion=2022-11-28
*/
type GitHubClient struct {
	Token string
	Url   string
}

func NewGithubClient(token string, repo *ProjectRepoInfo) *GitHubClient {
	var origin string = "https://api.github.com/repos"
	return &GitHubClient{
		Token: token,
		Url:   origin + "/" + repo.Owner + "/" + repo.Name + "/" + repo.ParseType,
	}
}

func (c *GitHubClient) check(res *http.Response) bool {
	links := res.Header.Get("Link")
	return links != ""
}

/*
해당 함수를 실행시킴으로 써 issue, pull request에 등록되어진 데이터를 조회합니다
입력 받는 함수로서 프로젝트의 owner와 repo 이름, 그리고 수집하고자 하는 데이터의 타입을 입력받습니다(issue,pr)
*/
func (c *GitHubClient) CallAPI(page int, ch chan<- []byte) {
	url := c.Url + fmt.Sprintf("?page=%d", page)
	resp, err := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Response was wrong, status code is %d", resp.StatusCode)
	}
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	ch <- data
}

func (c GitHubClient) Crawler() *chan []byte {
	ch := make(chan []byte)
	for i := 1; i < util.APICounter; i++ {
		go c.CallAPI(i, ch)

		issues := <-ch
		if len(issues) == 0 {
			break
		}
	}
	return &ch
}
