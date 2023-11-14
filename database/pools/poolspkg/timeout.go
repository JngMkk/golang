package poolspkg

import (
	"context"
	"time"
)

// 현재 시간을 얻어오려다 시간 초과를 발생시킴
func ExecWithTimeOut() error {
	db, err := Setup()
	if err != nil {
		return err
	}

	ctx := context.Background()

	// 즉시 시간 초과를 발생시킴
	ctx, cancel := context.WithDeadline(ctx, time.Now())

	defer cancel()

	// 트랜잭션이 컨텍스트를 인지함
	_, err = db.BeginTx(ctx, nil)

	return err
}
