package input

import "fineC/util"

type ProjectRepoInfo struct {
	Name      string
	Owner     string
	ParseType string
}

func NewProjectRepoInfo(name, owner string, ParseType int) *ProjectRepoInfo {
	var (
		ParseTypeStr string
	)

	switch ParseType {
	case util.Issues:
		ParseTypeStr = "issues"
	case util.PullRequest:
		ParseTypeStr = "pull"
	}

	return &ProjectRepoInfo{
		Name:      name,
		Owner:     owner,
		ParseType: ParseTypeStr,
	}
}
