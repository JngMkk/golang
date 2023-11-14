// close()로 채널을 닫아준다
// 좀비 고루틴(고루틴 릭) : 채널을 닫아주지 않아서 무한 대기를 하는 고루틴
package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Println("Square :", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch) // 채널을 닫아주어서 정상 종료되게 함
	wg.Wait()
}
