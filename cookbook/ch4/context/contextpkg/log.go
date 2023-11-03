package contextpkg

import (
	"context"

	"github.com/apex/log"
)

type key int

const logFields key = 0

func getFields(ctx context.Context) *log.Fields {
	fields, ok := ctx.Value(logFields).(*log.Fields)
	if !ok {
		f := make(log.Fields)
		fields = &f
	}
	return fields
}

// 항목과 context를 입력받고 context 객체로부터 데이터가 채워진 항목을 반환
func FromContext(ctx context.Context, l log.Interface) (context.Context, *log.Entry) {
	fields := getFields(ctx)
	e := l.WithFields(fields)
	ctx = context.WithValue(ctx, logFields, fields)
	return ctx, e
}

// context에 많은 로그 필드를 추가
func WithFields(ctx context.Context, fields log.Fielder) context.Context {
	f := getFields(ctx)
	for key, val := range fields.Fields() {
		(*f)[key] = val
	}
	ctx = context.WithValue(ctx, logFields, f)
	return ctx
}

// context에 로그 필드를 추가함
func WithField(ctx context.Context, key string, value interface{}) context.Context {
	return WithFields(ctx, log.Fields{key: value})
}
