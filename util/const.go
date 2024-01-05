package util

// ouptut으로 설정되어진 저장소로 저장하기위해 데이터 형태를 변환할 때 사용
const (
	DataToElasticsearch = 1 << iota
)

const (
	Issues = 10 << iota
	PullRequest
)
