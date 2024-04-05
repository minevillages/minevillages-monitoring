package db

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

type Redis struct {
	Addr     string
	Password string
	DBcount  int
}

type RedisSetter struct {
	Key    string
	Value  string
	Expire time.Duration
}

type RedisGetter struct {
	Key string
}

func (r *Redis) Connection() {
	// Connect to Redis server
	client = redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DBcount,
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

func (r *RedisSetter) Set() error {
	ctx := context.Background()
	return client.Set(ctx, r.Key, r.Value, r.Expire).Err()
}

func (r *RedisGetter) Get() (string, error) {
	ctx := context.Background()
	return client.Get(ctx, r.Key).Result()
}
