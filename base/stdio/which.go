package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}

	file := args[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path) // 운영체제에 따라 적절한 방법을 이용해 경로의 목록을 슬라이스로 분리해줌.

	for _, dir := range pathSplit {
		fullPath := filepath.Join(dir, file) // 운영체제마다 다른 구분자를 사용해 경로의 부분들을 합쳐줌.

		// 파일이 존재하는가?
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			// 일반 파일인가? (유닉스 환경에서는 모든 것이 파일이므로, 일반 파일도 실행 가능한 파일일 수 있음)
			if mode.IsRegular() {
				// 실행 파일인가?
				if mode&0111 != 0 {
					fmt.Println(fullPath)
				}
			}
		}
	}
}
