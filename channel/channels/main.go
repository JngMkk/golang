package main

import (
	"context"
	"time"

	"github.com/JngMkk/golang/channel/channels/channel"
)

func main() {
	ch := make(chan string)
	done := make(chan bool)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go channel.Printer(ctx, ch)
	go channel.Sender(ch, done)

	time.Sleep(2 * time.Second)
	done <- true
	cancel()

	time.Sleep(3 * time.Second)
}
