package healthapi

import (
	"context"
	"log"
	"minevillages/monitoring/api/db"
	"net/http"
	"time"
)

type HealthStatus struct {
	Status bool `json:"status"`
}

func MointorHealth() HealthStatus {
	healthStatus := HealthStatus{
		Status: true,
	}
	return healthStatus
}

type ServerInfo struct {
	Port string
	Name string
}

func (s *ServerInfo) Health() {
	url := "http://127.0.0.1:" + s.Port
	response, err := http.Get(url)

	if err != nil {
		log.Println(s.Name + " Server is unhealthy.")

		// Elastic Search 에 로그 저장
		doc := map[string]interface{}{
			"SERVER_URL": url,
		}
		client := db.EsClient
		_, err = client.Index().Index(s.Name).BodyJson(doc).Do(context.Background())
		if err != nil {
			log.Fatalf("Error indexing document: %v", err)
		}
		log.Println("Document indexed successfully to ElasticSearch")

		unhealthyGetter := &db.RedisGetter{
			Key: "unhealthy_" + s.Name,
		}
		_, err := unhealthyGetter.Get()

		if err != nil {
			unhealthySetter := &db.RedisSetter{
				Key:    "unhealthy_" + s.Name,
				Value:  url,
				Expire: 1 * time.Hour,
			}
			unhealthySetter.Set()
			mail := &Mail{
				Title: "[WARN] Server Unhealthy : " + s.Name,
				Body:  url,
			}
			mail.Send()
		}
		return
	}

	defer response.Body.Close()
}
