package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type requestContextKey struct{}
type requestContextValue struct {
	requestId string
}

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

			// logging requestId -> loggingProcessTimeMiddleware 바깥에 등록해야 요청에 id 부착 가능
			var requestId string
			ctx := r.Context()
			v := ctx.Value(requestContextKey{})

			if l, ok := v.(requestContextValue); ok {
				requestId = l.requestId
			}

			config.logger.Printf("requestId: %s, path:%s, method:%s, duration:%f\n",
				requestId, r.URL.Path, r.Method, time.Now().Sub(startTime).Seconds())
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

func addRequestIdMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// See part2-http-application/server/chapter5-http-server-tutorial/server-with-context/exercise.go
			requestId := strconv.FormatInt(time.Now().Unix(), 10)
			ctx := requestContextValue{
				requestId: requestId,
			}

			currentCtx := r.Context()
			newCtx := context.WithValue(currentCtx, requestContextKey{}, ctx)
			requestWithCtx := r.WithContext(newCtx)
			h.ServeHTTP(w, requestWithCtx)
		},
	)
}

func main() {
	config := appConfig{
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
	}

	mux := http.NewServeMux()
	registerHandler(mux, config)

	// middleware chaining addRequestIdMiddleware
	m := addRequestIdMiddleware(
		loggingProcessTimeMiddleware(
			panicHandlingMiddleware(mux, config),
			config,
		))

	log.Fatal(http.ListenAndServe(":7166", m))
}
