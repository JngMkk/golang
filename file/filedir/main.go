package main

import "github.com/JngMkk/golang/file/filedir/filedirs"

func main() {
	if err := filedirs.Operate(); err != nil {
		panic(err)
	}

	if err := filedirs.CaplitalizerExample(); err != nil {
		panic(err)
	}
}
