package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JngMkk/golang/network/dns/dnspkg"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <address>\n", os.Args[0])
		os.Exit(1)
	}

	addr := os.Args[1]
	lookup, err := dnspkg.LookupAddr(addr)
	if err != nil {
		log.Panicf("failed to lookup : %s", err.Error())
	}

	fmt.Println(lookup)
}
