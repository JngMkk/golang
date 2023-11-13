package main

import "fmt"

// 가변 인수 함수
func sum(nums ...int) int { // sum(nums []int) 랑 같음
	sum := 0

	// 슬라이스 타입
	fmt.Printf("nums 타입: %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20))
	fmt.Println(sum())
}
