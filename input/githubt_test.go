package input

import (
	"fineC/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client = GitHubClient{Token: util.NewToken()}

func TestParseData(t *testing.T) {
	test := assert.New(t)
	mockObj := NewProjectRepoInfo("python-mysql-replication", "julien-duponchelle", util.Issues)
	testObj := NewGithubClient(util.NewToken(), mockObj)
	result := testObj.CallApi()

	test.NotEmptyf(result, "Empty result")

}
