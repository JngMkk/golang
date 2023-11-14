// 함수 리터럴 ( 람다 )
package main

import "fmt"

type OpFn func(int, int) int

func getOperator(op string) OpFn {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	op := getOperator("+")

	var result = op(3, 4)
	fmt.Println(result)
}
