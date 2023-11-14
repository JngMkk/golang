package log

import (
	"log"

	"github.com/pkg/errors"
)

func OriginalError() error {
	return errors.New("error occurred")
}

// OriginalError 함수를 호출하고 래핑을 한 다음 오류 전달
func PassThroughError() error {
	err := OriginalError()
	// nil에도 동작하기 때문에 오류를 확인할 필요 없음
	return errors.Wrap(err, "in passthrougherror")
}

// 오류를 처리하며 전달은 하지 않음
func FinalDestination() {
	err := PassThroughError()
	if err != nil {
		log.Printf("an error occurred : %s\n", err.Error())
		return
	}
}
