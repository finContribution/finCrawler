package input

import (
	"fmt"
	"io"
	"net/http"
)

var url = "https://api.github.com/repos/"

/*
.env 파일에 추가한 token 정보를 가져옵니다 token 생성관련 내용은 링크 참조해주세요
https://docs.github.com/en/rest/authentication/authenticating-to-the-rest-api?apiversion=2022-11-28&apiVersion=2022-11-28
*/
type GitHubClient struct {
	Token string
	Repo  *ProjectRepoInfo
}

func NewGithubClient(token string, repo *ProjectRepoInfo) *GitHubClient {
	return &GitHubClient{
		token,
		repo,
	}
}

/*
해당 함수를 실행시킴으로 써 issue, pull request에 등록되어진 데이터를 조회합니다
입력 받는 함수로서 프로젝트의 owner와 repo 이름, 그리고 수집하고자 하는 데이터의 타입을 입력받습니다(issue,pr)
*/
func (c *GitHubClient) CallApi() []byte {

	url := fmt.Sprintf(url+"%s/%s/%s", c.Repo.Name, c.Repo.Owner, c.Repo.ParseType)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	req.Header.Set("Authorization", "token "+c.Token)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error fetching data:", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Non-OK HTTP status:", response.Status)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	return data
}
