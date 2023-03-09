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
	config appConfig

	// handler returns (status code, error)
	handler func(w http.ResponseWriter, r *http.Request, config appConfig) (int, error)
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// inject the config to the handler
	status, err := a.handler(w, r, a.config)
	if err != nil {
		a.config.logger.Printf("Response Status: %d, error: %s\n", status, err.Error())
		http.Error(w, err.Error(), status)
		return
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request, config appConfig) (int, error) {
	config.logger.Println("Handling API request")
	fmt.Fprintln(w, "This is API handler!")
	return http.StatusOK, nil
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request, config appConfig) (int, error) {
	if r.Method != http.MethodGet {
		return http.StatusMethodNotAllowed, fmt.Errorf("method not allowed: %s", r.Method)
	}

	config.logger.Println("Handling healthcheck request")
	fmt.Fprintln(w, "OK")
	return http.StatusOK, nil
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
