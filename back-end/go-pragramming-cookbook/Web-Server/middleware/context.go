package middleware

import (
	"context"
	"net/http"
	"strconv"
)

type ContextID string

const RequestID ContextID = "request-id"

func RequestIdMiddleware(start int64) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), RequestID, strconv.FormatInt(start, 10))

			// 요청 아이디 값 갱신
			start++

			// 값을 설정한 컨텍스트 객체를 다음 미들웨어의 요청 객체로 넘겨준다
			r = r.WithContext(ctx)

			next(w, r)
		}
	}
}

func GetRequestID(ctx context.Context) string {
	if val, ok := ctx.Value(RequestID).(string); ok {
		return val
	}
	return ""
}
