package statepkg

import (
	"context"
	"errors"
)

// 처리할 작업을 분기
func Processor(ctx context.Context, in chan *WorkRequest, out chan *WorkResponse) {
	for {
		select {
		case <-ctx.Done():
			return
		case wr := <-in:
			out <- Process(wr)
		}
	}
}

// 연산 타입별로 작업을 분기한 뒤 해당 연산 처리
func Process(wr *WorkRequest) *WorkResponse {
	resp := WorkResponse{Wr: wr}

	switch wr.Operation {
	case Add:
		resp.Result = wr.Value1 + wr.Value2
	case Sub:
		resp.Result = wr.Value1 - wr.Value2
	case Mul:
		resp.Result = wr.Value1 * wr.Value2
	case Div:
		if wr.Value2 == 0 {
			resp.Err = errors.New("divide by 0")
			break
		}
		resp.Result = wr.Value1 / wr.Value2
	default:
		resp.Err = errors.New("unsupported operation")
	}
	return &resp
}
