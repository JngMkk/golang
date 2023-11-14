package nulls

import (
	"encoding/json"
	"fmt"
)

type ExamplePointer struct {
	Age  *int   `json:"age,omitempty"`
	Name string `json:"name"`
}

// nil/omitted 값을 처리하는 방법
func PointerEncoding() error {
	e := ExamplePointer{}

	// no age는 nil age라는 점을 주의
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Pointer Unmarshal, no age: %+v\n", e)

	val, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Pointer Marshal, with no age:", string(val))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Pointer Unmarshal, with age = 0: %+v\n", e)

	val, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Pointer Marshal, with age = 0:", string(val))

	return nil
}
