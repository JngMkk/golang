package globalpkg

import (
	"errors"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

// 전역 패키지 수준 변수를 만듬
var (
	log     *logrus.Logger
	initLog sync.Once
)

// 처음에 로거를 설정하고 여러 번 실행되면 오류를 반환
func Init() error {
	err := errors.New("already initialized")
	initLog.Do(func() {
		err = nil
		log = logrus.New()
		log.Formatter = &logrus.JSONFormatter{}
		log.Out = os.Stdout
		log.Level = logrus.DebugLevel
	})
	return err
}

// 로그를 설정
func SetLog(l *logrus.Logger) {
	log = l
}

// 전역 로그에 WithField 함수가 연결된 로그를 내보냄
func WithField(key string, value interface{}) *logrus.Entry {
	return log.WithField(key, value)
}

// 전역 로그에 Debug 함수가 연결된 로그를 내보냄
func Debug(args ...interface{}) {
	log.Debug(args...)
}
