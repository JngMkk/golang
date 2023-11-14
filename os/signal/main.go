package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// SIGINT 인터럽트에 대한 리스너 설정
func CatchSig(ch chan os.Signal, done chan bool) {
	// 신호에 대기하기 위해 블록
	sig := <-ch

	// 신호를 받으면 출력
	fmt.Println("\nsig received :", sig)

	// 모든 신호에 대한 핸들러를 설정
	switch sig {
	case syscall.SIGINT:
		fmt.Println("handling a SIGINT now")
	case syscall.SIGTERM:
		fmt.Println("handling a SIGTERM in an entirely different way")
	default:
		fmt.Println("unexpected signal received")
	}

	// 종료
	done <- true
}

func main() {
	// 채널 초기화
	signals := make(chan os.Signal)
	done := make(chan bool)

	// signals 라이브러리 연결
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// 고루틴에 의해 신호가 잡혔으면 DONE
	go CatchSig(signals, done)
	fmt.Println("Press ctrl-c to terminate....")

	// 누군가 done으로 쓸 때까지 프로그램을 블록
	<-done
	fmt.Println("Done!")
}
