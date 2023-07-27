package main

import (
	"go.uber.org/zap"
	"net"
	"net/http"
	"time"
)

func WideEventLogMiddleware(logger *zap.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			wideWriter := &WideResponseWriter{ResponseWriter: w}
			start := time.Now()
			next.ServeHTTP(wideWriter, r)

			addr, _, _ := net.SplitHostPort(r.RemoteAddr)
			logger.Info("example wide event",
				zap.Int("status_code", wideWriter.status),
				zap.Int("response_length", wideWriter.length),
				zap.String("response_body", wideWriter.responseBody),
				zap.Int64("content_length", r.ContentLength),
				zap.String("method", r.Method),
				zap.String("proto", r.Proto),
				zap.String("remote_addr", addr),
				zap.String("uri", r.RequestURI),
				zap.String("user_agent", r.UserAgent()),
				zap.String("duration", time.Now().Sub(start).String()),
			)
		},
	)
}
