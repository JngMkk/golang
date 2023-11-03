package math

import (
	"fmt"
	"math"
)

func Examples() {
	i := 25

	res := math.Sqrt(float64(i))
	fmt.Println(res)

	res = math.Ceil(9.5)
	fmt.Println(res)

	fmt.Println("Pi :", math.Pi, "E :", math.E)
}
