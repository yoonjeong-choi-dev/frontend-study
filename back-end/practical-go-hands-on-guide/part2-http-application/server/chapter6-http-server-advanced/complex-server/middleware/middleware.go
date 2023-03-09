package middleware

import (
	"complex-server/config"
	"fmt"
	"net/http"
	"time"
)

func loggingProcessTimeMiddleware(h http.Handler, conf config.AppConfig) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			h.ServeHTTP(w, r)
			conf.Logger.Printf("protocol:%s, path:%s, method:%s, duration:%f\n",
				r.Proto, r.URL.Path, r.Method, time.Now().Sub(startTime).Seconds())
		},
	)
}

func panicHandlingMiddleware(h http.Handler, conf config.AppConfig) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Check the panic and handling
			defer func() {
				if rValue := recover(); rValue != nil {
					conf.Logger.Println("panic detected", rValue)
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintf(w, "unexpected server error")
				}
			}()

			h.ServeHTTP(w, r)
		},
	)
}
