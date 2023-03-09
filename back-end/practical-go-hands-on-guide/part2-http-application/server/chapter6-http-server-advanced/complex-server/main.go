package main

import (
	"complex-server/config"
	"complex-server/handlers"
	"complex-server/middleware"
	"io"
	"log"
	"net/http"
	"os"
)

func setupServer(mux *http.ServeMux, w io.Writer) http.Handler {
	conf := config.InitConfig(w)
	handlers.RegisterHandler(mux, conf)
	return middleware.RegisterMiddleware(mux, conf)
}

func main() {
	mux := http.NewServeMux()
	wrappedMux := setupServer(mux, os.Stdout)

	log.Fatal(http.ListenAndServe(":7166", wrappedMux))
}
