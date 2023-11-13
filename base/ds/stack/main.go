package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	}
	return nil
}

func main() {
	s := NewStack()
	books := [5]string{"어린왕자", "겨울왕국", "노인과 바다", "빨간머리 앤", "짱구"}

	for i := 0; i < 5; i++ {
		s.Push(books[i])
	}

	val := s.Pop()
	for val != nil {
		fmt.Printf("%v -> ", val)
		val = s.Pop()
	}

	fmt.Println()
}
