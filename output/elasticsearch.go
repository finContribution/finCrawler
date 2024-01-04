package output

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"log"
	"strings"
)

// MyDocument - Elasticsearch에 인덱싱할 데이터 구조체
type MyDocument struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// IndexDocumentsToElasticsearch - 채널에서 받은 데이터를 Elasticsearch에 인덱싱하는 함수
func IndexDocumentsToElasticsearch(dataChan <-chan MyDocument, es *elasticsearch.Client, indexName string) {
	// 벌크 인덱서 설정
	bulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:  indexName, // 사용할 인덱스 이름
		Client: es,        // Elasticsearch 클라이언트
	})
	if err != nil {
		log.Fatalf("Error creating the bulk indexer: %s", err)
	}

	// 채널에서 데이터 읽기 및 벌크 인덱싱
	for doc := range dataChan {
		data, err := json.Marshal(doc)
		if err != nil {
			log.Fatalf("Error marshaling document: %s", err)
		}

		err = bulkIndexer.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: doc.ID,
				Body:       strings.NewReader(string(data)),
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, resp esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						log.Printf("ERROR: %s", err)
					} else {
						log.Printf("ERROR: %s: %s", resp.Error.Type, resp.Error.Reason)
					}
				},
			},
		)
		if err != nil {
			log.Fatalf("Error adding item to bulk indexer: %s", err)
		}
	}

	// 벌크 인덱싱 완료 확인
	if err := bulkIndexer.Close(context.Background()); err != nil {
		log.Fatalf("Error closing bulk indexer: %s", err)
	}

	// 인덱싱 상태 확인
	stats := bulkIndexer.Stats()
	log.Printf("Indexed documents: %d, failed: %d", stats.NumIndexed, stats.NumFailed)
}
