package runner

import (
	"fineC/input"
	"fineC/output"
	"fineC/util"
	"github.com/elastic/go-elasticsearch/v7"
)

type Runner struct {
	Input    input.Input
	Output   output.Output
	URL      string
	SaveHost string
}

func NewRunner(inputFlag, outputFlag, url string) *Runner {
	var inputObject input.Input
	var outputObject output.Output

	switch inputFlag {
	case "github":
		inputObject = &input.GitHubClient{Token: util.NewToken("../.env"), Url: url}

	}
	switch outputFlag {
	case "elasticsearch":
		client, _ := elasticsearch.NewDefaultClient()
		outputObject = &output.ElasticsearchOutput{
			Client:    client,
			IndexName: "test-index",
		}
	}

	return &Runner{
		Input:  inputObject,
		Output: outputObject,
	}
}

func (r *Runner) Run() {
	ch := make(chan []byte, util.APICounter)
	r.Input.Crawling(ch)
	r.Output.Convert(ch)
	r.Output.Send()
}
