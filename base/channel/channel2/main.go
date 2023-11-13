// 채널 기본 크기 : 0
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go square()

	// 데이터를 넣었는데 가져가는 애가 없으므로 계속 대기
	ch <- 9
	fmt.Println("Never print")
}

func square() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("sleep")
	}
}
