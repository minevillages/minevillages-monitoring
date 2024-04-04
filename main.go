package main

import (
	"log"
	"minevillages/monitoring/api"
	baseapi "minevillages/monitoring/api/base"
	"minevillages/monitoring/api/db"
	healthapi "minevillages/monitoring/api/health"
	"net/http"
	"runtime"

	"github.com/joho/godotenv"
	// "github.com/quic-go/quic-go/http3"
)

func main() {
	// 프로세스의 모든 CPU 코어를 활용하도록 설정합니다.
	runtime.GOMAXPROCS(runtime.NumCPU())

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	handler := api.HostRouteHandler{}
	handler.RegisterHost("127.0.0.1:8080", baseapi.HTTPHandler{})

	db.RedisConnection()
	db.ElasticSearchConnection()
	healthapi.TimeSchedule()

	if err := http.ListenAndServe(
		":8080",
		// "ssl/certificate.crt",
		// "ssl/private.key",
		&handler,
	); err != nil {
		log.Fatalln("HTTP 관련 수신 소켓을 초기화하는 과정에서 예외가 발생하였습니다.", err.Error())
	}

}
