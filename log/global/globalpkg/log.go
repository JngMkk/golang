package globalpkg

// 전역 로그를 사용하는 방법을 보여줌
func UseLog() error {
	if err := Init(); err != nil {
		return err
	}

	// 다른 패키지 안에 있는 경우,
	// 이 값은 globalpkg.WithField와 globalpkg.Debug를 통해 가져올 수 있음
	WithField("key", "value").Debug("hello")
	Debug("test")
	Debug("test2")

	return nil
}
