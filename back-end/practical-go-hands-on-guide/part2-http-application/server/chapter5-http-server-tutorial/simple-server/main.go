package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Exercise 5.1
type requestLogData struct {
	URL           string `json:"url"`
	Method        string `json:"method"`
	ContentLength int64  `json:"content_length"`
	Protocol      string `json:"protocol"`
}

func logRequest(req *http.Request) {
	logData := requestLogData{
		URL:           req.URL.String(),
		Method:        req.Method,
		ContentLength: req.ContentLength,
		Protocol:      req.Proto,
	}

	data, err := json.Marshal(&logData)
	if err != nil {
		panic(err)
	}

	log.Println(string(data))
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	fmt.Fprintf(w, "Hello, this is a simple server")
}

func heathCheckHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	fmt.Fprintf(w, "OK")
}

func registerHandler(mux *http.ServeMux) {
	mux.HandleFunc("/api", apiHandler)
	mux.HandleFunc("/health", heathCheckHandler)
}

func main() {
	listenArr := os.Getenv("LISTEN_ADDR")
	if len(listenArr) == 0 {
		listenArr = ":7166"
	}

	mux := http.NewServeMux()
	registerHandler(mux)

	log.Fatal(http.ListenAndServe(listenArr, mux))
}
