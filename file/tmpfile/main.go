package main

import "github.com/JngMkk/golang/file/tmpfile/tmpFiles"

func main() {
	if err := tmpFiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
