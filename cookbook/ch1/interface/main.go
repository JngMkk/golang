package main

import (
	"bytes"
	"fmt"

	"github.com/JngMkk/cookbook/ch1/interface/interfaces"
)

func main() {
	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}
	fmt.Print("Stdout on Copy = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}

	fmt.Println("out bytes buffer =", out.String())
	fmt.Print("Stdout on PipeExample = ")
	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}
}
