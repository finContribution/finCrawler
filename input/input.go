package input

type Input interface {
	CallAPI(interface{}) []byte
	Crawling() chan []byte
}
