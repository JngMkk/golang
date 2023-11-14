package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// http 연결을 가져와 웹소켓으로 변환
// 권장 버퍼 크기 사용
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// 연결 업그레이드
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("failed to upgrade connection :", err)
		return
	}

	for {
		// 루프에서 메시지를 읽고 동일한 메시지를 전달함
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("failed to read message :", err)
			return
		}
		log.Printf("received from client : %#v", string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println("failed to write message :", err)
			return
		}
	}
}

func main() {
	fmt.Println("Listening on port : 8000")

	// 모든 요청을 처리하기 위해 localhost:8000에 하나의 핸들러 마운트
	log.Panic(http.ListenAndServe("localhost:8000", http.HandlerFunc(wsHandler)))
}
