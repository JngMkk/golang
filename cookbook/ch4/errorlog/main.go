package main

import (
	"fmt"

	"github.com/JngMkk/cookbook/ch4/errorlog/log"
)

func main() {
	fmt.Println("basic logging and modification of logger : ")
	log.Log()
	fmt.Println("logging 'handled' errors : ")
	log.FinalDestination()
}
