package middleware

import (
	"complex-server/config"
	"net/http"
)

func RegisterMiddleware(mux *http.ServeMux, conf config.AppConfig) http.Handler {
	return loggingProcessTimeMiddleware(
		panicHandlingMiddleware(mux, conf),
		conf)
}
