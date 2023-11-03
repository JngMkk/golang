package main

import (
	"github.com/JngMkk/cookbook/ch6/storage/storagepkg"
)

func main() {
	if err := storagepkg.Exec(); err != nil {
		panic(err)
	}
}
