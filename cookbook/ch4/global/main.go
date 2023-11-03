package main

import "github.com/JngMkk/cookbook/ch4/global/globalpkg"

func main() {
	if err := globalpkg.UseLog(); err != nil {
		panic(err)
	}
}
