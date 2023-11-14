package main

import (
	"fmt"

	"github.com/JngMkk/golang/error/errorwrap/errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()

	errwrap.Unwrap()
	fmt.Println()

	errwrap.StackTrace()
}
