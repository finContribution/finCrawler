package input

type Input interface {
	Crawling() chan [][]byte
}
