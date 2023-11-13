package main

import (
	"base/pkg/banking"
	"fmt"
)

func main() {
	account := banking.Account{Owner: "eddie", Balance: 1000}
	fmt.Println(account)
}
