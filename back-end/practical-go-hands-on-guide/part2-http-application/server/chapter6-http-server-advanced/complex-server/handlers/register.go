package handlers

import (
	"complex-server/config"
	"net/http"
)

func RegisterHandler(mux *http.ServeMux, conf config.AppConfig) {
	mux.Handle("/api", &app{conf: conf, handler: apiHandler})
	mux.Handle("/health", &app{conf: conf, handler: healthCheckHandler})
	mux.Handle("/panic", &app{conf: conf, handler: panicHandler})
}
