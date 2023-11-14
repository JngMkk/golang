package main

import "github.com/JngMkk/golang/database/pools/poolspkg"

func main() {
	if err := poolspkg.ExecWithTimeOut(); err != nil {
		panic(err)
	}
}
