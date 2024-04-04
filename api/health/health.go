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

func ServerHealth(port string, name string) {
	url := "http://127.0.0.1:" + port
	response, err := http.Get(url)

	if err != nil {
		log.Println(name + " Server is unhealthy.")

		// Elastic Search 에 로그 저장
		doc := map[string]interface{}{
			"SERVER_URL": "http://127.0.0.1:" + port,
		}

		client := db.EsClient
		_, err = client.Index().Index(name).BodyJson(doc).Do(context.Background())
		if err != nil {
			log.Fatalf("Error indexing document: %v", err)
		}
		log.Println("Document indexed successfully to ElasticSearch")

		_, err := db.GetRedis("unhealthy_" + name)
		if err != nil {
			db.SetRedis("unhealthy_"+name, port, 1*time.Hour)
			SendMail("[WARN] Server Unhealthy : "+name, "http://127.0.0.1:"+port)

		}
		return
	}

	defer response.Body.Close()
}
