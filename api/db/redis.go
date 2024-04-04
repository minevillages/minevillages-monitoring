package db

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func RedisConnection() {
	// Connect to Redis server
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // Redis server address
		Password: "",               // No password
		DB:       0,                // Default DB
	})

	// Ping Redis to ensure connectivity
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Println("Failed to ping Redis:", err)
		return
	} else {
		log.Println("Success to ping Redis.")
	}
}

func SetRedis(key, value string, expire time.Duration) error {
	ctx := context.Background()
	return client.Set(ctx, key, value, expire).Err()
}

func GetRedis(key string) (string, error) {
	ctx := context.Background()
	return client.Get(ctx, key).Result()
}
