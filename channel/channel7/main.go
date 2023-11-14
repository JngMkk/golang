// 컨텍스트
// 작업을 지시할 때 작업 가능 시간, 작업 취소 등의 조건을 지시할 수 있는 작업 명세서 역할

// 작업 취소 컨텍스트
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())

	// 작업 시간 설정
	// ctx, cancel = context.WithTimeout(context.Background(), 3 * time.Second) -> 3초 뒤에 ctx.Done() 채널에 시그널 발생

	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)

	cancel()
	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Println("tick")
		}
	}
}
