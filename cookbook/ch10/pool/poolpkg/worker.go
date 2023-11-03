package poolpkg

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type op string

type WorkRequest struct {
	Op      op
	Text    []byte
	Compare []byte
}

type WorkResponse struct {
	Wr      WorkRequest
	Result  []byte
	Matched bool
	Err     error
}

const (
	// bcrypt 작업 타입
	Hash op = "encrypt"
	// bcrypt 비교 작업
	Compare op = "decrypt"
)

func hashWork(wr WorkRequest) WorkResponse {
	val, err := bcrypt.GenerateFromPassword(wr.Text, bcrypt.DefaultCost)
	return WorkResponse{
		Result: val,
		Err:    err,
		Wr:     wr,
	}
}

func compareWork(wr WorkRequest) WorkResponse {
	var matched bool
	err := bcrypt.CompareHashAndPassword(wr.Compare, wr.Text)
	if err == nil {
		matched = true
	}
	return WorkResponse{
		Matched: matched,
		Err:     err,
		Wr:      wr,
	}
}

// 워커 풀 채널에 작업을 전달
func Process(wr WorkRequest) WorkResponse {
	switch wr.Op {
	case Hash:
		return hashWork(wr)
	case Compare:
		return compareWork(wr)
	default:
		return WorkResponse{Err: errors.New("unsupported operation")}
	}
}

// 계속 반복 실행되며 워커 풀의 일부
func Worker(ctx context.Context, id int, in chan WorkRequest, out chan WorkResponse) {
	for {
		select {
		case <-ctx.Done():
			return
		case wr := <-in:
			fmt.Printf("worker id: %d, performing %s worker\n", id, wr.Op)
			out <- Process(wr)
		}
	}
}

// 매개변수로 전달된 numWorker 수만큼의 워커를 생성하고
// 작업 및 응답 추가를 위한 함수와 cancel 함수를 반환함 cancel 함수는 반드시 호출되어야 함
func Dispatch(numWorker int) (context.CancelFunc, chan WorkRequest, chan WorkResponse) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	in := make(chan WorkRequest, 10)
	out := make(chan WorkResponse, 10)

	for i := 0; i < numWorker; i++ {
		go Worker(ctx, i, in, out)
	}
	return cancel, in, out
}
