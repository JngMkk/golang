package interfaces

import (
	"io"
	"os"
)

func PipeExample() error {
	// pipe reader와 pipe writer는 io.Reader와 io.Writer를 구현함
	r, w := io.Pipe()

	// reader가 정리를 위해 마지막에 닫을 때까지 대기하기 위해 블록 방식으로 동작하기 때문에 별도의 go 루틴에서 실행해야 함
	go func() {
		w.Write([]byte("test\n"))
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}

	return nil
}
