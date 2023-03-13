package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handleAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Response: TLS Server")
}

func registerHandlersAndMiddlewares(mux *http.ServeMux, logger *log.Logger) http.Handler {
	mux.HandleFunc("/api", handleAPI)
	return loggingMiddleware(mux, logger)
}

func loggingMiddleware(h http.Handler, logger *log.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			h.ServeHTTP(w, r)
			logger.Printf(
				"protocol: %s, path: %s, method: %s, duration: %f\n",
				r.Proto, r.URL.Path, r.Method, time.Now().Sub(startTime).Seconds(),
			)
		},
	)
}

func main() {
	tlsCertFile := os.Getenv("TLS_CERT_FILE_PATH")
	tlsKeyFile := os.Getenv("TLS_KEY_FILE_PATH")

	if len(tlsCertFile) == 0 || len(tlsKeyFile) == 0 {
		log.Fatal("TLS_CERT_FILE_PATH and TLS_KEY_FILE_PATH must both be specified in environment variable")
	}

	mux := http.NewServeMux()
	logger := log.New(
		os.Stdout, "tls-server",
		log.Lshortfile|log.LstdFlags,
	)

	mainHandler := registerHandlersAndMiddlewares(mux, logger)
	log.Fatal(
		http.ListenAndServeTLS(
			":7166", tlsCertFile, tlsKeyFile, mainHandler,
		),
	)
}
