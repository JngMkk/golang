package main

import (
	"fmt"

	"github.com/JngMkk/cookbook/ch3/tag/tags"
)

func main() {
	if err := tags.EmptyStruct(); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := tags.FullStruct(); err != nil {
		panic(err)
	}
}
