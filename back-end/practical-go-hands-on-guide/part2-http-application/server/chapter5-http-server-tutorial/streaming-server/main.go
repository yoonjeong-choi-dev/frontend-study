package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/decode", decodeLogStreamHandler)
	mux.HandleFunc("/streaming", streamingHandler)

	http.ListenAndServe(":7166", mux)
}
