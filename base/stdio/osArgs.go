package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}

	var min, max float64
	for i := 1; i < len(args); i++ {
		n, err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			continue
		}
		if i == 1 {
			min = n
			max = n
			continue
		}
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}

/*
go run base/stdio/osArgs.go a b 5 1

Min:  0
Max:  5
*/
