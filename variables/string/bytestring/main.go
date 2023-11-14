package main

import "github.com/JngMkk/golang/variables/string/bytestring/bytestrings"

func main() {
	err := bytestrings.WorkWithBuffer()
	if err != nil {
		panic(err)
	}

	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()
}
