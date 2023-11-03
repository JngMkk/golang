package errwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

// 오류 래핑과 오류에 주석을 표시
func WrappedError(e error) error {
	return errors.Wrap(e, "An error occurred in WrappedError")
}

type ErrorTyped struct {
	error
}

// 오류를 래핑할 때 호출
func Wrap() {
	e := errors.New("standard errors")
	fmt.Println("Regular Error -", WrappedError(e))
	fmt.Println("Typed Error -", WrappedError(ErrorTyped{errors.New("typed error")}))
	fmt.Println("Nil -", WrappedError(nil))
}
