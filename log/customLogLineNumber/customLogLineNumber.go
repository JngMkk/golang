package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "log.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Lshortfile: 소스코드의 파일명과 줄 번호를 로그에 추가
	LstdFlags := log.Ldate | log.Lshortfile
	logger := log.New(f, "logger ", LstdFlags)
	logger.Println("Hello, there!")

	logger.SetFlags(log.Lshortfile | log.LstdFlags)
	logger.Println("Another log entry!")
}

/*
logger 2023/11/14 customLogLineNumber.go:22: Hello, there!
logger 2023/11/14 15:58:32 customLogLineNumber.go:25: Another log entry!
*/
