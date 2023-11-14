// 외부에서 로직을 주입하는 것을 의존성 주입이라고 함
package main

import (
	"fmt"
	"os"
)

type Writer func(string)

// type WriterInterface interface {
// 	Write(string)
// }

// 어떤 역할을 할지 모름(외부에서 주입)
func writeHello(writer Writer) {
	writer("Hello World")
}

// func writeHello2(writer WriterInterface) {
// 	writer.Write("Hello World")
// }

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a File")
		return
	}
	defer f.Close()

	writeHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}
