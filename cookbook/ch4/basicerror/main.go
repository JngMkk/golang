package main

import (
	"fmt"

	"github.com/JngMkk/cookbook/ch4/basicerror/basicerrors"
)

func main() {
	basicerrors.BasicErrors()
	err := basicerrors.SomeFunc()
	fmt.Println("custom error :", err)
}
