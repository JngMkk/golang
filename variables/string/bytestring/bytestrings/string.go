package bytestrings

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// 문자열을 검색하는 여러 방법을 보여줌
func SearchString() {
	s := "this is a test"

	// 문자열이 단어 this를 포함하기 때문에 true
	fmt.Println(strings.Contains(s, "this"))

	// 문자열이 철자 a를 포함하기 때문에 true
	fmt.Println(strings.ContainsAny(s, "abc"))

	// 문자열이 this로 시작하기 때문에 true
	fmt.Println(strings.HasPrefix(s, "this"))

	// 문자열이 test로 끝나기 때문에 true
	fmt.Println(strings.HasSuffix(s, "test"))
}

// 문자열을 수정하는 방법을 보여줌
func ModifyString() {
	s := "simple string"

	// simple과 string을 출력
	fmt.Println(strings.Split(s, " "))

	s = " simple string "

	// 앞뒤 공백을 제거
	fmt.Println(strings.TrimSpace(s))
}

// 문자열로 io.Reader 인터페이스를 빠르게 생성하는 방법
func StringReader() {
	s := "simple string"
	r := strings.NewReader(s)

	// 표준 출력에 출력
	io.Copy(os.Stdout, r)
}
