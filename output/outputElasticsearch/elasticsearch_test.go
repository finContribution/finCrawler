package outputElasticsearch

import (
	"reflect"
	"testing"
	"time"
)

// TestParseChan 함수는 ParseChan 메소드를 테스트합니다.
func TestParseChan(t *testing.T) {
	// Mock 데이터 생성
	jsonData := []byte(`[{"name": "Alice", "age": 30}]`)
	inputChan := make(chan []byte, 1)
	outputChan := make(chan map[string]interface{}, 1)

	eo := ElasticsearchOutput{OutputDataStream: outputChan}
	go eo.Convert(inputChan)

	// Input 채널에 데이터 전달
	inputChan <- jsonData
	close(inputChan) // 채널 닫기

	// 결과를 기다리는 타임아웃 설정
	select {
	case result := <-outputChan:
		expected := map[string]interface{}{"name": "Alice", "age": float64(30)} // JSON 파싱 결과는 float64로 처리될 수 있음

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Got %v, want %v", result, expected)
		}

	case <-time.After(2 *
		time.Second): // 1초 타임아웃
		t.Error("Test timed out")
	}
}
