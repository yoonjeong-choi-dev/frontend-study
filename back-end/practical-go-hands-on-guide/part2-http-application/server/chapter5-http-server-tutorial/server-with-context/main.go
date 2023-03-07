package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type requestContextKey struct{}
type requestContextValue struct {
	requestId string
}

func addRequestId(r *http.Request, requestId string) *http.Request {
	ctx := requestContextValue{
		requestId: requestId,
	}

	currentCtx := r.Context()
	// WithValue returns a copy of parent in which the value associated with key is val
	newCtx := context.WithValue(currentCtx, requestContextKey{}, ctx)

	return r.WithContext(newCtx)
}

func createRequestId(r *http.Request) string {
	return fmt.Sprintf("[%s]%s-%s", time.Now().String(), r.Method, r.URL.String())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	nextReq := addRequestId(r, createRequestId(r))
	processRequest(w, nextReq)
}

func processRequest(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	fmt.Fprint(w, "Request Processed Successful")
}

func logRequest(r *http.Request) {
	ctx := r.Context()
	v := ctx.Value(requestContextKey{})

	if l, ok := v.(requestContextValue); ok {
		log.Printf("Processing request Id: %s\n", l.requestId)
	}
}

func main() {
	listenArr := os.Getenv("LISTEN_ADDR")
	if len(listenArr) == 0 {
		listenArr = ":7166"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	log.Fatal(http.ListenAndServe(listenArr, mux))
}
