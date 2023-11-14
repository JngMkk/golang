package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var total, nInts, nFloats int
	args := os.Args
	invalid := make([]string, 0)

	for _, val := range args[1:] {
		// 값이 정수인가?
		_, err := strconv.Atoi(val)
		if err == nil {
			total++
			nInts++
			continue
		}

		// 부동소수점인가?
		_, err = strconv.ParseFloat(val, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}

		// 유효한 숫자 x
		invalid = append(invalid, val)
	}

	fmt.Printf("#read: %d #ints: %d #floats: %d #invalid: %d\n", len(args), nInts, nFloats, len(invalid))
}

/*
go run base/stdio/osArgs2.go 1 2 3 3.1 a b 5.1

#read: 8 #ints: 3 #floats: 2 #invalid: 2
*/
