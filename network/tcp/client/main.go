package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const addr = "localhost:8888"

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// 클라이언트로부터의 문자열 입력을 가져옴
		fmt.Printf("Enter some text : ")
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("encountered an error reading input : %s\n", err.Error())

			continue
		}

		// addr에 연결
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Printf("encountered an error connecting : %s\n", err.Error())
		}

		// 연결에 데이터를 작성
		fmt.Fprintln(conn, data)

		// 응답을 다시 읽어옴
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("encountered an error reading response : %s\n", err.Error())
		}
		fmt.Printf("Received back : %s\n", status)
		conn.Close()
	}
}
