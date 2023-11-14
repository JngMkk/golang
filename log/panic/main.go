package main

import (
	"fmt"

	"github.com/JngMkk/golang/log/panic/panicpkg"
)

func main() {
	fmt.Println("before panic")
	panicpkg.Catcher()
	fmt.Println("after panic")
}
