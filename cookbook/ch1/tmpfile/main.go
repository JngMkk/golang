package main

import "github.com/JngMkk/cookbook/ch1/tmpfile/tmpfiles"

func main() {
	if err := tmpfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
