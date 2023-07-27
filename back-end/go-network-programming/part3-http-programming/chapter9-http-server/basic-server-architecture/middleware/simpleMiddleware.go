package middleware

import (
	"log"
	"net/http"
	"time"
)

func SimpleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Function 1. Filter HTTP Method
			if r.Method == http.MethodTrace {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			// Function 2. Add a header
			w.Header().Set("X-Content-Type-Options", "nosiff")

			// Function 3. Logging
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Printf("Next handler duration: %v\n", time.Now().Sub(start))
		})
}
