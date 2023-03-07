package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type appConfig struct {
	logger *log.Logger
}

// app : custom http.Handler implementation
type app struct {
	config  appConfig
	handler func(w http.ResponseWriter, r *http.Request, config appConfig)
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// inject the config to the handler
	a.handler(w, r, a.config)
}

func apiHandler(w http.ResponseWriter, r *http.Request, config appConfig) {
	config.logger.Println("Handling API request")
	fmt.Fprintln(w, "This is API handler!")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request, config appConfig) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	config.logger.Println("Handling healthcheck request")
	fmt.Fprintln(w, "OK")
}

func panicHandler(w http.ResponseWriter, r *http.Request, config appConfig) {
	panic("server will be dead")
}

func registerHandler(mux *http.ServeMux, config appConfig) {
	mux.Handle("/api", &app{config: config, handler: apiHandler})
	mux.Handle("/health", &app{config: config, handler: healthCheckHandler})
	mux.Handle("/panic", &app{config: config, handler: panicHandler})
}

func loggingProcessTimeMiddleware(h http.Handler, config appConfig) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			h.ServeHTTP(w, r)
			config.logger.Printf("path:%s, method:%s, duration:%f\n",
				r.URL.Path, r.Method, time.Now().Sub(startTime).Seconds())
		},
	)
}

func panicHandlingMiddleware(h http.Handler, config appConfig) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Check the panic and handling
			defer func() {
				if rValue := recover(); rValue != nil {
					config.logger.Println("panic detected", rValue)
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintf(w, "unexpected server error")
				}
			}()

			h.ServeHTTP(w, r)
		},
	)
}

func main() {
	config := appConfig{
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
	}

	mux := http.NewServeMux()
	registerHandler(mux, config)

	// middleware chaining
	m := loggingProcessTimeMiddleware(panicHandlingMiddleware(mux, config), config)

	log.Fatal(http.ListenAndServe(":7166", m))
}
