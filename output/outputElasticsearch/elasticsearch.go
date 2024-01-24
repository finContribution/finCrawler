package outputElasticsearch

import (
	"context"
	"encoding/json"
	"fineC/util"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"log"
	"strings"
)

type ElasticsearchOutput struct {
	Context          context.Context
	Client           *elasticsearch.Client
	IndexName        string
	OutputDataStream chan map[string]interface{}
}

// TODO 데이터 convert 후 'OutputDataStream' 채널에 대한 관리를 어떻게 할 것인지 파악 필요
func (eo *ElasticsearchOutput) Convert(input chan []byte) {
	eo.OutputDataStream = make(chan map[string]interface{}, util.APICounter*10)
	for data := range input {
		go func(data []byte) { // 고루틴으로 병렬 처리
			var parseDatas []map[string]interface{}
			err := json.Unmarshal(data, &parseDatas)
			if err != nil {
				log.Printf("Can't insert to outputElasticsearch: %v", err) // 에러 로깅 변경
				return                                                     // 에러 발생 시 다음 데이터로 넘어감
			}
			for _, parseData := range parseDatas {
				eo.OutputDataStream <- parseData
			}
		}(data)
	}
}

// TODO 모든 데이터 전송완료 후 시스템의 중단 시점을 찾아야함 Context를 통한 채널 닫는 것으로 생각 중
// 생성된 인스턴스의 OutputDataStream을 통해서 converted 된 데이터를 처리하기 위한 로직입니다.
func (eo *ElasticsearchOutput) Send() {
	// 벌크 인덱서 설정
	bulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:  eo.IndexName, // 사용할 인덱스 이름
		Client: eo.Client,    // Elasticsearch 클라이언트
	})
	if err != nil {
		log.Fatalf("Error creating the bulk indexer: %s", err)
	}

	// 슬라이스에서 데이터 읽기 및 벌크 인덱싱
	for data := range eo.OutputDataStream {
		convertedData, err := json.Marshal(data)
		if err != nil {
			log.Panicf("Somthing Wrong: %d", err)
		}
		err = bulkIndexer.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action: "index",
				Body:   strings.NewReader(string(convertedData)),
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, resp esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						log.Printf("ERROR: %s", err)
					} else {
						log.Printf("ERROR: %s: %s", resp.Error.Type, resp.Error.Reason)
					}
				},
			},
		)
	}

	if err != nil {
		log.Fatalf("Error adding item to bulk indexer: %s", err)
	}

	// 벌크 인덱싱 완료 확인
	if err := bulkIndexer.Close(context.Background()); err != nil {
		log.Fatalf("Error closing bulk indexer: %s", err)
	}

	// 인덱싱 상태 확인
	stats := bulkIndexer.Stats()
	log.Printf("Indexed documents: %d, failed: %d", stats.NumIndexed, stats.NumFailed)
}
