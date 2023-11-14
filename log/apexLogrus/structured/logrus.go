package structured

import "github.com/sirupsen/logrus"

type Hook struct {
	id string
}

// 로그를 기록할 때마다 호출
func (hook *Hook) Fire(entry *logrus.Entry) error {
	entry.Data["id"] = hook.id
	return nil
}

// 전달된 hook이 실행될 레벨
func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// 몇 가지 기본적인 logrus 기능을 보여줌
func Logrus() {
	// json 포맷으로 로그를 기록
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.AddHook(&Hook{"123"})

	fields := logrus.Fields{}
	fields["success"] = true
	fields["complex_struct"] = struct {
		Event string
		When  string
	}{"Somethig happend", "Just Now"}

	x := logrus.WithFields(fields)
	x.Warn("warning")
	x.Error("error")
}
