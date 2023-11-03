package main

import "github.com/JngMkk/cookbook/ch1/filedir/filedirs"

func main() {
	if err := filedirs.Operate(); err != nil {
		panic(err)
	}

	if err := filedirs.CaplitalizerExample(); err != nil {
		panic(err)
	}
}
