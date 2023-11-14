package main

import (
	"context"
	"fmt"

	"github.com/JngMkk/golang/channel/state/statepkg"
)

func main() {
	in := make(chan *statepkg.WorkRequest, 10)
	out := make(chan *statepkg.WorkResponse, 10)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go statepkg.Processor(ctx, in, out)
	req := statepkg.WorkRequest{statepkg.Add, 3, 4}
	in <- &req

	req2 := statepkg.WorkRequest{statepkg.Sub, 5, 2}
	in <- &req2

	req3 := statepkg.WorkRequest{statepkg.Mul, 9, 9}
	in <- &req3

	req4 := statepkg.WorkRequest{statepkg.Div, 8, 2}
	in <- &req4

	req5 := statepkg.WorkRequest{statepkg.Div, 8, 0}
	in <- &req5

	for i := 0; i < 5; i++ {
		resp := <-out
		fmt.Printf("Request: %v, Result: %v, Error: %v\n", resp.Wr, resp.Result, resp.Err)
	}
}
