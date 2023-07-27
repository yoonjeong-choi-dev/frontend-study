package main

import (
	"net/http"
)

type WideResponseWriter struct {
	http.ResponseWriter
	length       int
	status       int
	responseBody string
}

func (w *WideResponseWriter) WriteHeader(status int) {
	w.ResponseWriter.WriteHeader(status)
	w.status = status
}

func (w *WideResponseWriter) Write(b []byte) (int, error) {
	n, err := w.ResponseWriter.Write(b)
	w.length += n

	if w.status == 0 {
		w.status = http.StatusOK
	}
	if err == nil {
		w.responseBody = string(b[:n])
	}
	return n, err
}
