package output

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestElasticOutput_Convert(t *testing.T) {
	test_obj := []byte(`[
    {"id": 1, "name": "Alice", "email": "alice@example.com", "isActive": true, "scores": [85, 92, 78], "metadata": {"role": "admin", "department": "sales"}},
    {"id": 2, "name": "Bob", "email": "bob@example.com", "isActive": false, "scores": [72, 88, 91], "metadata": {"role": "user", "department": "tech"}}
]`)
	client, _ := elasticsearch.NewDefaultClient()
	eo := ElasticOutput{Client: client, Schema: "test-index"}
	obj := eo.Convert(test_obj)
	test := assert.New(t)

	var expected []map[string]interface{}
	json.Unmarshal(test_obj, &expected)
	test.Equal(expected, obj)

}

func TestElasticOutput_Receive(t *testing.T) {
	ch := make(chan []byte, 2)
	client, _ := elasticsearch.NewDefaultClient()
	eo := ElasticOutput{Client: client, Schema: "test-index"}
	test_obj := []byte(`[
    {"id": 1, "name": "Alice", "email": "alice@example.com", "isActive": true, "scores": [85, 92, 78], "metadata": {"role": "admin", "department": "sales"}},
    {"id": 2, "name": "Bob", "email": "bob@example.com", "isActive": false, "scores": [72, 88, 91], "metadata": {"role": "user", "department": "tech"}}
]`)
	eo.Receive(ch)
	ch <- test_obj
	close(ch)

	for data := range eo.ConvertedData {
		fmt.Println(data)
	}

}
