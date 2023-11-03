package main

import "github.com/JngMkk/cookbook/ch6/pools/poolspkg"

func main() {
	if err := poolspkg.ExecWithTimeOut(); err != nil {
		panic(err)
	}
}
