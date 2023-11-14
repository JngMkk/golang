package main

import "fmt"

func main() {
	// 일반적인 for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()

	// 관용적 for loop
	j := 0
	for ok := true; ok; ok = (j != 10) {
		fmt.Print(j*j, " ")
		j++
	}
	fmt.Println()

	// while loop
	k := 0
	for {
		if k == 10 {
			break
		}
		fmt.Print(k*k, " ")
		k++
	}
	fmt.Println()
}
