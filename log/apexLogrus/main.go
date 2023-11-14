package main

import (
	"fmt"

	"github.com/JngMkk/golang/log/apexLogrus/structured"
)

func main() {
	fmt.Println("Logrus :")
	structured.Logrus()

	fmt.Println()
	fmt.Println("Apex :")
	structured.Apex()
}
