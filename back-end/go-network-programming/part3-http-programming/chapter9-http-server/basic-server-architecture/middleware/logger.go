package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Function 3. Logging
			log.Printf("Request: %s\n", r.URL.Path)
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Printf("Next handler duration: %v\n", time.Now().Sub(start))
		})
}
