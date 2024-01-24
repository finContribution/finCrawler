package inputGithub

import (
	"fineC/util"
	"fmt"
	"io"
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

func NewGithubClient(token string, repo string) *GitHubClient {
	var origin string = "https://api.github.com/repos/"
	return &GitHubClient{
		Token: token,
		Url:   origin + repo + "/issues",
	}
}

/*
해당 함수를 실행시킴으로 써 issue, pull request에 등록되어진 데이터를 조회합니다
입력 받는 함수로서 프로젝트의 owner와 repo 이름, 그리고 수집하고자 하는 데이터의 타입을 입력받습니다(issue,pr)
*/
func (c *GitHubClient) CallAPI(page int) ([]byte, error) {
	url := c.Url + fmt.Sprintf("?page=%d", page)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("token %s", c.Token))
	resp, err := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		message, _ := io.ReadAll(resp.Body)
		err := fmt.Errorf("Response was wrong, status code is %d, message: %s", resp.StatusCode, message)
		return nil, err

	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, err
	}
	return data, nil
}

/*
CallAPI 호출을 통해 불러 온 모든 데이터를 채널에 적재할 수 있도록 합니다
(페이징 된 데이터를 지속적으로 불러옵니다. 대신 APICounter에 등록되어진 요청할 수 있는 한계를 지정합니다)
channel이 close 되지 않은 상태임에 따라 사용 시 주의가 필요함
*/
func (c GitHubClient) Crawling(ch chan []byte) {
	for i := 1; i < util.APICounter; i++ {
		data, err := c.CallAPI(i)
		if err != nil {
			panic(err)
		}
		ch <- data
	}
	close(ch)
}
