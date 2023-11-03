package main

import "github.com/JngMkk/cookbook/ch1/bytestring/bytestrings"

func main() {
	err := bytestrings.WorkWithBuffer()
	if err != nil {
		panic(err)
	}

	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()
}
