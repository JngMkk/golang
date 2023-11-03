package main

import (
	"fmt"

	"github.com/JngMkk/cookbook/ch4/errorwrap/errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()

	errwrap.Unwrap()
	fmt.Println()

	errwrap.StackTrace()
}
