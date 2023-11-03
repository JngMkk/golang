package structured

import (
	"errors"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

// 추적할 오류 발생
func ThrowError() error {
	err := errors.New("a crazy failure")
	log.WithField("id", "123").Trace("ThrowError").Stop(&err)
	return err
}

// 두 개의 스트림으로 분할
type CustomHandler struct {
	id      string
	handler log.Handler
}

// hook을 더해 로그를 기록
func (h *CustomHandler) HandleLog(e *log.Entry) error {
	e.WithField("id", h.id)
	return h.handler.HandleLog(e)
}

func Apex() {
	log.SetHandler(&CustomHandler{"123", text.New(os.Stdout)})
	err := ThrowError()

	log.WithError(err).Error("an error occured")
}
