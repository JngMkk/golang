package main

import (
	"fmt"

	"github.com/JngMkk/golang/log/errorlog/log"
)

func main() {
	fmt.Println("basic logging and modification of logger : ")
	log.Log()
	fmt.Println("logging 'handled' errors : ")
	log.FinalDestination()
}
