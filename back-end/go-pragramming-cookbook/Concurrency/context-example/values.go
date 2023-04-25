package context_example

import "context"

// key context 객체에 저장할 key-value 데이터의 키 타입
type key string

// context 객체에 저장시 사용할 키 정의
const (
	timeoutKey  key = "TimeoutKey"
	deadlineKey key = "DeadlineKey"
)

func Setup(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, timeoutKey, "timeout exceeded")
	ctx = context.WithValue(ctx, deadlineKey, "deadline exceeded")
	return ctx
}

func GetValueFromContext(ctx context.Context, k key) string {
	if val, ok := ctx.Value(k).(string); ok {
		return val
	}
	return ""
}
