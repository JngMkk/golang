package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/JngMkk/cookbook/ch2/envconfig/envvar"
)

type Config struct {
	Version string `json:"version" required:"true"`
	IsSafe  bool   `json:"is_safe" default:"true"`
	Secret  string `json:"secret"`
}

func main() {
	// 예제 json 파일을 저장하기 위해 임시 파일 생성
	tf, err := os.CreateTemp("", "tmp")
	if err != nil {
		panic(err)
	}
	defer tf.Close()
	// defer os.Remove(tf.Name())

	// secrets 값을 저장하기 위해 json 파일 생성
	secrets := `{
		"secret": "so so secret"	
	}`

	if _, err = tf.Write(bytes.NewBufferString(secrets).Bytes()); err != nil {
		panic(err)
	}

	// 필요한 경우 환경 변수를 쉽게 설정할 수 있음
	if err = os.Setenv("EXAMPLE_VERSION", "1.0.0"); err != nil {
		panic(err)
	}

	if err = os.Setenv("EXAMPLE_ISSAFE", "false"); err != nil {
		panic(err)
	}

	c := Config{}
	if err = envvar.LoadConfig(tf.Name(), "EXAMPLE", &c); err != nil {
		panic(err)
	}

	fmt.Println("secrets file contains =", secrets)

	// 환경 변수를 쉽게 읽을 수 있음
	fmt.Println("EXAMPLE_VERSION =", os.Getenv("EXAMPLE_VERSION"))
	fmt.Println("EXAMPLE_ISSAFE =", os.Getenv("EXAMPLE_ISSAFE"))

	// 최종 구성은 json과 환경 변수가 혼합돼 있음
	fmt.Printf("Final Config: %#v\n", c)
}
