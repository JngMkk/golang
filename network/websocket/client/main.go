package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/gorilla/websocket"
)

// ctrl-c를 막기 때문에 종료하려면 ctrl-c를 누르고 엔터 키를 누르거나 다른 위치에서 종료
func process(c *websocket.Conn) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Enter some text : ")

		data, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read stdin", err)
		}

		// 읽어온 문자열에서 공백 제거
		data = strings.TrimSpace(data)

		// 웹소켓에 메시지를 바이트로 작성
		err = c.WriteMessage(websocket.TextMessage, []byte(data))
		if err != nil {
			log.Println("failed to write message :", err)
			return
		}

		// 에코 서버이므로 작성한 후에 읽어올 수 있음
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("failed to read :", err)
			return
		}
		log.Printf("received back from server : %#v\n", string(message))
	}
}

// ctrl-c로 프로그램을 종료하면 웹소켓 연결을 정리함
func catchSig(ch chan os.Signal, c *websocket.Conn) {
	// os.Signal을 대기하기 위해 블록
	<-ch

	err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("write close :", err)
	}
}

func main() {
	// 채널에 os.Signal 값 연결
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// 웹소켓에 연결하기 위해 ws:// 구조 사용
	u := "ws://localhost:8000/"
	log.Printf("connecting to %s", u)

	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal("dial :", err)
	}
	defer c.Close()

	// catchSig에 전달
	go catchSig(interrupt, c)

	process(c)
}
