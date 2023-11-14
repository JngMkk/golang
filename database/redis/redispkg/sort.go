package redispkg

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

func Sort() error {
	ctx := context.Background()

	conn, err := Setup()
	if err != nil {
		return err
	}

	listkey := "list"
	if err := conn.LPush(ctx, listkey, 1).Err(); err != nil {
		return err
	}
	// 다음 명령 중 하나라도 오류가 발생하면 리스트 키를 제거함
	defer conn.Del(ctx, listkey)

	if err := conn.LPush(ctx, listkey, 3).Err(); err != nil {
		return err
	}

	if err := conn.LPush(ctx, listkey, 2).Err(); err != nil {
		return err
	}

	res, err := conn.Sort(ctx, listkey, &redis.Sort{Order: "ASC"}).Result()
	if err != nil {
		return err
	}
	fmt.Println(res)

	return nil
}