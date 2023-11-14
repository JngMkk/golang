package main

import (
	"context"
	"fmt"

	"github.com/JngMkk/golang/channel/pipeline/pipelinepkg"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	in, out := pipelinepkg.NewPipeline(ctx, 10, 2)

	go func() {
		for i := 0; i < 20; i++ {
			in <- fmt.Sprint("Message", i)
		}
	}()

	for i := 0; i < 20; i++ {
		<-out
	}
}
