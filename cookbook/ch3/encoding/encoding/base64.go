package encoding

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

// base64 패키지를 사용하는 데모
func Base64Example() error {
	// base64 패키지는 바이트/문자열로 동작하는 바이너리 포맷을 지원할 수 없을 때 유용
	value := base64.URLEncoding.EncodeToString([]byte("encoding some data"))
	fmt.Println("With EncodeToString and URLEncoding: ", value)

	decoded, err := base64.URLEncoding.DecodeString(value)
	if err != nil {
		return err
	}
	fmt.Println("With DecodeString and URLEncoding: ", string(decoded))

	return nil
}

func Base64ExampleEncoder() error {
	buf := bytes.Buffer{}

	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	if _, err := encoder.Write([]byte("encoding some other data")); err != nil {
		return err
	}

	if err := encoder.Close(); err != nil {
		return err
	}
	fmt.Println("Using encoder and StdEncoding: ", buf.String())

	decoder := base64.NewDecoder(base64.StdEncoding, &buf)
	res, err := ioutil.ReadAll(decoder)
	if err != nil {
		return err
	}
	fmt.Println("Using decoder and StdEncoding: ", string(res))

	return nil
}
