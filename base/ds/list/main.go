package main

import (
	"container/list"
	"fmt"
)

// 삽입 : O(1)
// 인덱스 요소 접근 : O(N)

// 데이터 지역성
// 데이터가 인접해 있을수록 캐시 성공률이 올라가고 성능도 증가함
// 일반적으로 요소 수가 적은 경우 리스트보다 배열이 빠름
func main() {
	v := list.New()
	e4 := v.PushBack(4)
	e1 := v.PushFront(1)
	v.InsertBefore(3, e4)
	v.InsertAfter(2, e1)

	// 1 2 3 4
	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

	// 4 3 2 1
	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}
