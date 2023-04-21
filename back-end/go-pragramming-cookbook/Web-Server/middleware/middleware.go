package middleware

import (
	"log"
	"net/http"
	"time"
)

// Middleware 미들웨어로 구현되는 함수 타입 정의
// => ApplyMiddlewares 함수에서 일괄적으로 미들웨어를 등록하기 위한 타입 정의
// Chapter 7. Web Client 의 client-wrapping-middleware/middleware.go 와 동일한 방식
type Middleware func(http.HandlerFunc) http.HandlerFunc

func ApplyMiddlewares(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	applied := h

	// middlewares 앞 부분이 들어오는 요청에 대해 가장 나중에 처리
	for _, middleware := range middlewares {
		applied = middleware(applied)
	}
	return applied
}

func Logger(l *log.Logger) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			l.Printf("started request to %s with request-id %s\n", r.URL, GetRequestID(r.Context()))

			next(w, r)

			l.Printf("completed request to %s with request-id %s in %s\n",
				r.URL, GetRequestID(r.Context()),
				time.Since(start),
			)
		}
	}
}
