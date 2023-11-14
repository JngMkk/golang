package main

import "github.com/JngMkk/golang/log/global/globalpkg"

func main() {
	if err := globalpkg.UseLog(); err != nil {
		panic(err)
	}
}
