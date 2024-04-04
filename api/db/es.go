package db

import (
	"log"

	"github.com/olivere/elastic/v7"
)

var EsClient *elastic.Client

func ElasticSearchConnection() {
	esURL := "http://127.0.0.1:9200"
	client, err := elastic.NewClient(elastic.SetURL(esURL))
	if err != nil {
		log.Fatalf("Failed to create Elasticsearch client: %v", err)
	}
	EsClient = client
}
