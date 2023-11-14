package main

import (
	"fmt"

	"github.com/JngMkk/golang/channel/pool/poolpkg"
)

func main() {
	cancel, in, out := poolpkg.Dispatch(10)
	defer cancel()

	for i := 0; i < 10; i++ {
		in <- poolpkg.WorkRequest{Op: poolpkg.Hash, Text: []byte(fmt.Sprintf("messages %d", i))}
	}

	for i := 0; i < 10; i++ {
		res := <-out
		if res.Err != nil {
			panic(res.Err)
		}
		in <- poolpkg.WorkRequest{Op: poolpkg.Compare, Text: res.Wr.Text, Compare: res.Result}
	}

	for i := 0; i < 10; i++ {
		res := <-out
		if res.Err != nil {
			panic(res.Err)
		}
		fmt.Printf("string: \"%s\", matched: %v\n", string(res.Wr.Text), res.Matched)
	}
}
