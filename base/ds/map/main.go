// golang은 hashmap 사용 (unordered)
// map의 원리 (hash)
// 같은 입력이 들어오면 같은 결과가 나온다
// 다른 입력이 들어오면 되도록 다른 결과가 나온다
// 입력값의 범위는 무한대이고 결과는 특정 범위를 갖는다
// 파일의 변조 체크 가능 (해킹 방지)
package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["화랑"] = "서울시 광진구"
	m["하나"] = "서울시 강남구"
	m["유현"] = "부산시 사하구"
	m["다인"] = "대구시 중구"
	m["다현"] = "서울시 금천구"

	fmt.Printf("하나의 주소는 %s입니다.\n", m["하나"])
	fmt.Printf("다인의 주소는 %s입니다.\n", m["다인"])
}
