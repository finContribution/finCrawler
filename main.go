package main

import (
	"fineC/input"
	"fineC/input/inputGithub"
	elasticsearch2 "fineC/output/outputElasticsearch"
	runner2 "fineC/runner"
	"fineC/util"
	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	repo := input.NewProjectRepoInfo("kubernetes", "kubernetes", util.Issues)
	input := inputGithub.NewGithubClient(util.NewToken("./.env"), repo)
	client, _ := elasticsearch.NewDefaultClient()
	output := &elasticsearch2.ElasticsearchOutput{Client: client, IndexName: "test-index"}

	runner := &runner2.Runner{
		Input:  input,
		Output: output,
	}

	runner.Run()
	//for data := range output.OutputDataStream {
	//	fmt.Println(data)
	//}

}
