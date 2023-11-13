// 값이 없을 때 기본값
package main

import "fmt"

func main() {
	m := make(map[int]int)

	m[1] = 0
	m[2] = 2
	m[3] = 3

	delete(m, 3)
	fmt.Println(m[3]) // 값이 있어서 0인지 없어서 0인지 모름
	fmt.Println(m[1])

	v, ok := m[3]
	fmt.Println(v, ok)
	v, ok = m[1]
	fmt.Println(v, ok)

	if v, ok = m[3]; ok {
		fmt.Println("있어요", v)
	} else {
		fmt.Println("없어요")
	}
}
