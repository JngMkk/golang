// 일반함수는 상태를 가질 수 없지만 함수 리터럴은 내부 상태를 가질 수 있음.
// 캡처는 값 복사가 아닌 레퍼런스 복사(포인터가 복사)
package main

import "fmt"

func main() {
	i := 0

	f := func() {
		i += 10
	}

	i++

	f()

	fmt.Println(i)
}
