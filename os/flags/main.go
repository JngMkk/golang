package main

import (
	"flag"
	"fmt"
)

func main() {
	// 설정 초기화
	c := Config{}
	c.Setup()

	//  일반적으로 main에서 호출됨
	flag.Parse()

	fmt.Println(c.GetMessage())
}
