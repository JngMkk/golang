package main

import (
	"fmt"

	"github.com/JngMkk/golang/error/basicerror/basicerrors"
)

func main() {
	basicerrors.BasicErrors()
	err := basicerrors.SomeFunc()
	fmt.Println("custom error :", err)
}
