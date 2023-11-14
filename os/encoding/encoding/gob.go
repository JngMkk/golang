package encoding

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type pos struct {
	X      int
	Y      int
	Object string
}

// gob 패키지를 사용하는 방법
func GobExample() error {
	buf := bytes.Buffer{}

	p := pos{
		X:      10,
		Y:      15,
		Object: "wrench",
	}

	// p가 인터페이스인 경우에는 먼저 gob.Register을 호출해야 함

	e := gob.NewEncoder(&buf)
	if err := e.Encode(&p); err != nil {
		return err
	}
	fmt.Println("Gob Encoded valued length: ", len(buf.Bytes()))
	fmt.Println("Gob Encoded value: ", buf.Bytes())

	p2 := pos{}
	d := gob.NewDecoder(&buf)
	if err := d.Decode(&p2); err != nil {
		return err
	}
	fmt.Println("Gob Decode value: ", p2)

	return nil
}
