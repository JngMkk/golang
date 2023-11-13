// 동시성 프로그래밍 주의
// 동일한 메모리 자원을 여러 고루틴에서 접근할 때 동시성 문제가 발생
package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
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

	// account.Balance에 동시 접근하고 있음
	// 메모리 자원을 하나의 고루틴에서만 접근하게 해야 함
	// Lock을 사용하는 것 (뮤텍스 Mutual Exclusion)
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
