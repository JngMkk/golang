package nulls

import (
	"encoding/json"
	"fmt"
)

const (
	jsonBlob     = `{"name": "Aaron"}`
	fulljsonBlob = `{"name": "Aaron", "age": 0}`
)

type Example struct {
	Age  int    `json:"age,omitempty"`
	Name string `json:"name"`
}

// 일반적인 타입에 대한 인코딩 및 디코딩 방법
func BaseEncoding() error {
	e := Example{}

	// no age가 0 age라는 점에 주의
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshal, no age: %+v\n", e)

	val, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Regular Marshal, with no age:", string(val))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshal, with age = 0: %+v\n", e)

	val, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Regular Marshal, with age = 0:", string(val))

	return nil
}
