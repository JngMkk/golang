package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 파일을 입력받아 각 단어를 키로 하고 단어의 등장 수를 값으로 하는 맵을 반환
func WordCount(f io.Reader) map[string]int {
	res := make(map[string]int)

	// io.Reader 인터페이스에서 동작할 스캐너를 만듬
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		res[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return res
}

func main() {
	fmt.Printf("string: number_of_occurrences\n\n")

	for k, v := range WordCount(os.Stdin) {
		fmt.Printf("%s: %d\n", k, v)
	}
}
