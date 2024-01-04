package input

import (
	"fineC/util"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

var client = GitHubClient{Token: util.NewToken()}
var wg sync.WaitGroup

func TestParseData(t *testing.T) {
	test := assert.New(t)
	repoName := "python-mysql-replication"
	testObject := client.ParseData("julien-duponchelle", repoName, "issues")

	test.NotEmptyf(testObject, "Empty result")

}
