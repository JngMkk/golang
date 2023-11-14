package main

import (
	"github.com/JngMkk/golang/database/storage/storagepkg"
)

func main() {
	if err := storagepkg.Exec(); err != nil {
		panic(err)
	}
}
