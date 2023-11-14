package panicpkg

import (
	"fmt"
	"strconv"
)

// 0으로 나누기로 인한 패닉 발생
func Panic() {
	zero, err := strconv.ParseInt("0", 10, 64)
	if err != nil {
		panic(err)
	}
	a := 1 / zero
	fmt.Println("we'll never get here", a)
}

// panic 함수 호출
func Catcher() {
	// 대부분의 웹 애플리케이션에서는 패닉이 발생하면 해당 패닉을 포착해
	// http.InternalServerErrror 메시지를 표시하는 것이 일반적임.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic occurred :", r)
		}
	}()
	Panic()
}
