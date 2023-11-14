package redispkg

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func Exec() error {
	ctx := context.Background()

	conn, err := Setup()
	if err != nil {
		return err
	}

	c1 := "value"
	conn.Set(ctx, "key", c1, 5*time.Second)

	var result string
	if err := conn.Get(ctx, "key").Scan(&result); err != nil {
		switch err {
		//키가 발견되지 않았다는 것을 의미
		case redis.Nil:
			return nil
		default:
			return err
		}
	}
	fmt.Println("result =", result)

	return nil
}
