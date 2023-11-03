package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file", err)
		return
	}
	defer fmt.Println("반드시 호출됨")
	defer f.Close()
	defer fmt.Println("파일을 닫았음")

	fmt.Println("파일에 Hello World를 씁니다")
	fmt.Fprintln(f, "Hello World")
}
