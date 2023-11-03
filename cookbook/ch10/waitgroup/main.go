package main

import (
	"fmt"

	"github.com/JngMkk/cookbook/ch10/waitgroup/waitgrouppkg"
)

func main() {
	sites := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://www.google.com/search?q=golang",
	}

	resps, err := waitgrouppkg.Crawl(sites)
	if err != nil {
		panic(err)
	}
	fmt.Println("Resps received:", resps)
}
