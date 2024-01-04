package output

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"testing"
)

func TestIndexDocumentsToElasticsearch(t *testing.T) {
	// Elasticsearch 클라이언트 설정 (예제에서는 기본 클라이언트 사용)
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 데이터 채널 생성 및 함수 호출 예제
	dataChan := make(chan MyDocument)
	go func() {
		dataChan <- MyDocument{ID: "1", Name: "John Doe"}
		dataChan <- MyDocument{ID: "2", Name: "Jane Doe"}
		close(dataChan)
	}()

	IndexDocumentsToElasticsearch(dataChan, es, "test-index")
}
