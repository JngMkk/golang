package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "customLog.log")

	// os.OpenFile()을 호출하면 작성할 로그 파일이 만들어짐. 파일이 이미 있다면 파일을 열어 파일의 끝에 새 데이터를 추가함(os.O_APPEND)
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	// defer 키워드와 함께 작성한 구문은 현재 함수가 끝나기 전에 실행됨.
	defer f.Close()

	logger := log.New(f, "logger ", log.LstdFlags)
	logger.Println("Hello there!")
	logger.Println("----------------------------")
}

/*
logger 2023/11/14 15:59:41 Hello there!
logger 2023/11/14 15:59:41 ----------------------------
*/
