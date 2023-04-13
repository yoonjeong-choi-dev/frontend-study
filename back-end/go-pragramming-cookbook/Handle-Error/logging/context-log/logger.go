package context_log

import (
	"context"
	"github.com/apex/log"
)

type IdKey int

// logFields context 로깅 키
const logFields IdKey = 7166

func GetFields(ctx context.Context) *log.Fields {
	fields, ok := ctx.Value(logFields).(*log.Fields)
	if !ok {
		// context 객체에 log.Fields 객체가 없는 경우 새로 추가
		f := make(log.Fields)
		fields = &f
	}
	return fields
}

// FromContext context 객체와 로그 객체에 서로 공유하는 log.Fields 객체 연결
func FromContext(ctx context.Context, l log.Interface) (context.Context, *log.Entry) {
	fields := GetFields(ctx)
	e := l.WithFields(fields)
	ctx = context.WithValue(ctx, logFields, fields)
	return ctx, e
}

// WithField context 객체에 로그 필드 추가
func WithField(ctx context.Context, key string, value interface{}) context.Context {
	return WithFields(ctx, log.Fields{key: value})
}

// WithFields context 객체에 로그 필드 추가
func WithFields(ctx context.Context, fields log.Fielder) context.Context {
	f := GetFields(ctx)
	for key, val := range fields.Fields() {
		(*f)[key] = val
	}
	return context.WithValue(ctx, logFields, f)
}
