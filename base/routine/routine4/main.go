// 뮤텍스의 문제점
// 동시성 프로그래밍으로 인한 성능 향상을 얻을 수 없음
// 과도한 락킹으로 성능이 하락되기도 함
// 고루틴을 완전히 멈추게 만드는 데드락 문제 발생
// 영역을 나누는 방법 & 역할을 나누는 방법으로 해결 가능
package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

var mutex sync.Mutex

func DepositAndWithdraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()

	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Microsecond)
	account.Balance -= 1000
}

func main() {
	var wg sync.WaitGroup

	account := &Account{10}
	wg.Add(10)

	// 한 개의 루틴만 접근하므로 프로그램 종료 x
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
