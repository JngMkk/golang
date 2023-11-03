package contextpkg

import (
	"context"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

func gatherName(ctx context.Context) {
	ctx = WithField(ctx, "name", "Go Cookbook")
}

func gatherLocation(ctx context.Context) {
	ctx = WithFields(ctx, log.Fields{"city": "Seattle", "state": "WA"})
}

// 설정을 위해 세 개의 함수 호출을 하고 종료 전에 로그 기록
func Initialize() {
	// 기본 로그 설정
	log.SetHandler(text.New(os.Stdout))

	// context 초기화
	ctx := context.Background()

	// 로거를 생성하고 context에 연결
	ctx, e := FromContext(ctx, log.Log)

	// 필드 설정
	ctx = WithField(ctx, "id", "123")
	e.Info("starting")
	gatherName(ctx)
	e.Info("after gatherName")
	gatherLocation(ctx)
	e.Info("after gatherLocation")
}
