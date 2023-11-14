package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type connections struct {
	addrs map[string]*net.UDPAddr
	// 맵의 수정을 위해 락시킴
	mu sync.Mutex
}

func broadcast(conn *net.UDPConn, conns *connections) {
	count := 0

	for {
		count++
		conns.mu.Lock()

		// 확인한 주소에 대해 루프로 반복 처리
		for _, retAddr := range conns.addrs {
			// 모두에게 메시지 전송
			msg := fmt.Sprintf("Sent %d", count)
			if _, err := conn.WriteToUDP([]byte(msg), retAddr); err != nil {
				fmt.Printf("error encountered : %s\n", err.Error())
				continue
			}
		}

		conns.mu.Unlock()
		time.Sleep(1 * time.Second)
	}
}
