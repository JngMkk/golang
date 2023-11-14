package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongo 클라이언트 초기화
func Setup(ctx context.Context, addr string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	// cancel은 설정이 종료될 때 호출
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(addr))
	if err != nil {
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	return client, nil
}
