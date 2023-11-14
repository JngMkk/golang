/*
Go의 동시성 모델은 고루틴과 채널을 이용해 구현함.

goroutine은 Go에서 실행할 수 있는 가장 작은 단위.
새로운 고루틴을 만들고 싶으면 go 키워드를 사용하고 그 뒤에 미리 정의된 함수나 익명 함수를 호출하면 됨.

Channel은 Go에서 고루틴끼리 통신하고 데이터를 주고받을 수 있게 하는 메커니즘.

고루틴을 만들기는 쉽지만 동시성 프로그래밍에는 고루틴을 동기화하고 데이터를 공유하는 등의 다른 어려움이 있음.
고루틴을 실행할 때 부수 효과가 발생하는 것을 피해야 하기 때문.
main()도 고루틴의 하나로 실행되고 main() 고루틴의 실행이 끝나면 아직 끝나지 않은 고루틴을 포함한 전체 프로그램이 종료되기 때문에
다른 고루틴들이 끝나기 전에 main()의 실행이 끝나지 않도록 주의해야 함.

고루틴끼리는 변수를 공유하지 않지만 메모리는 공유할 수 있음.
main()에서 고루틴들이 채널이나 메모리를 통해 데이터를 주고받는 것을 기다릴 수 있게 하는 여러 가지 기법이 있음.
	ex) time.Sleep()
*/

package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	go PrintHangul()
	go PrintNumbers()

	time.Sleep(3 * time.Second)
}
