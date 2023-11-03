package errwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

// 오류의 래핑을 해제하고 오류의 타입 assertion 수행
func Unwrap() {
	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")
	fmt.Println("wrapped error :", err)

	// 다양한 오류 타입 처리 가능
	switch errors.Cause(err).(type) {
	case ErrorTyped:
		fmt.Println("a typed error occurred :", err)
	default:
		fmt.Println("an unknown error occurred")
	}
}

// 오류에 대한 모든 스택 출력
func StackTrace() {
	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")

	fmt.Printf("%+v\n", err)
}
