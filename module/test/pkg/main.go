package main

import (
	"fmt"

	"github.com/JngMkk/golang/module/test/pkg/banking"
)

func main() {
	account := banking.Account{Owner: "eddie", Balance: 1000}
	fmt.Println(account)
}
