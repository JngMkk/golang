// 패닉 : 처리하기 힘든 에러를 만났을 때 프로그램을 조기 종료하는 방법
// 빠르게 종료시켜서 오류를 해결하기 위해서 사용
package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		panic("b는 0일 수 없습니다")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	divide(9, 3)
	divide(9, 0)
}
