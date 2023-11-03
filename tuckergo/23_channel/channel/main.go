// 채널은 고루틴 간의 메세지 큐
// Thread-safe queue
// 멀티스레드 환경에서 Lock 없이 쓸 수 있음
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)
	ch <- 9
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch

	time.Sleep(time.Second)
	fmt.Println("Square :", n*n)
	wg.Done()
}
