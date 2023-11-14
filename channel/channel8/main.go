// 컨텍스트 특정 값 설정
package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	// 키, 값
	ctx := context.WithValue(context.Background(), "number", 9)
	go square(ctx)

	wg.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Square : %d\n", n*n)
	}
	wg.Done()
}
