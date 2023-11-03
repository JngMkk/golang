package redispkg

import (
	"context"

	"github.com/go-redis/redis"
)

// redis 클라이언트 초기화
func Setup() (*redis.Client, error) {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		// default DB
		DB: 0,
	})

	_, err := client.Ping(ctx).Result()
	return client, err
}
