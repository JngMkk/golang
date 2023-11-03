package main

import (
	"fmt"

	"github.com/JngMkk/cookbook/ch4/panic/panicpkg"
)

func main() {
	fmt.Println("before panic")
	panicpkg.Catcher()
	fmt.Println("after panic")
}
