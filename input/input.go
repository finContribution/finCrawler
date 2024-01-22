package input

type Input interface {
	CallAPI(page int) ([]byte, error)
	Crawling(ch chan []byte)
}
