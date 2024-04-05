package db

import (
	"log"

	"github.com/olivere/elastic/v7"
)

var EsClient *elastic.Client

type ElasticSearch struct {
	ESUrl string
}

func (e *ElasticSearch) Connection() {

	client, err := elastic.NewClient(elastic.SetURL(e.ESUrl))
	if err != nil {
		log.Fatalf("Failed to create Elasticsearch client: %v", err)
	}
	EsClient = client
}
