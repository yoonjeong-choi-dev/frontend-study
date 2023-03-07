package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

func registerHandler(mux *http.ServeMux, config appConfig) {
	mux.Handle("/api", &app{config: config, handler: apiHandler})
	mux.Handle("/health", &app{config: config, handler: healthCheckHandler})
}

func main() {
	config := appConfig{
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
	}

	mux := http.NewServeMux()
	registerHandler(mux, config)

	log.Fatal(http.ListenAndServe(":7166", mux))
}
