package dataconv

import (
	"fmt"
	"strconv"
)

// 일부 strconv 예제를 보여줌
func Strconv() error {
	// strconv는 문자열에서 다른 타입으로 또는 다른 타입에서 문자열로 변환하는 좋은 방법
	s := "1234"

	// 진수 설정에 10을 지정하고, 비트 사이즈에 64비트를 지정
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	fmt.Println(res)

	// 16진수 변환
	res, err = strconv.ParseInt("FF", 16, 64)
	if err != nil {
		return err
	}

	fmt.Println(res)

	val, err := strconv.ParseBool("true")
	if err != nil {
		return err
	}

	fmt.Println(val)

	return nil
}
