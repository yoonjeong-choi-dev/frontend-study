package handlers

import (
	"complex-server/config"
	"fmt"
	"net/http"
)

type app struct {
	conf    config.AppConfig
	handler func(w http.ResponseWriter, r *http.Request, conf config.AppConfig)
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// inject the config to the handler
	a.handler(w, r, a.conf)
}

func apiHandler(w http.ResponseWriter, r *http.Request, conf config.AppConfig) {
	conf.Logger.Println("Handling API request")
	fmt.Fprint(w, "This is API handler!")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request, conf config.AppConfig) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	conf.Logger.Println("Handling healthcheck request")
	fmt.Fprint(w, "OK")
}

func panicHandler(w http.ResponseWriter, r *http.Request, conf config.AppConfig) {
	panic("server will be dead")
}
