package healthapi

import (
	"time"
)

func TimeSchedule() {
	// 3초마다 health check
	ticker := time.NewTicker(3 * time.Second)

	for range ticker.C {
		server := &ServerInfo{
			Port: "4000/api/user/health",
			Name: "user",
		}
		server.Health()
	}

}
