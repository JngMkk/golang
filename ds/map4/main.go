// 해쉬로 맵 만들기
// 해쉬 충돌 주의 (리스트로 키,값 넣기로 해결 가능 )
package main

import "fmt"

const M = 10

func hash(d int) int {
	return d % M
}

func main() {
	m := [M]string{}

	m[hash(23)] = "하나"
	m[hash(259)] = "유현"

	fmt.Printf("%d = %s\n", 23, m[hash(23)])
	fmt.Printf("%d = %s\n", 259, m[hash(259)])
}
