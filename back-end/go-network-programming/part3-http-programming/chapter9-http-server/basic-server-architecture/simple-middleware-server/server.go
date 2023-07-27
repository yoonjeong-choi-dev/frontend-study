package main

import (
	"basic-server-architecture/middleware"
	"fmt"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi~"))
}

func main() {
	mux := http.NewServeMux()
	simpleMiddleware := middleware.SimpleMiddleware(http.HandlerFunc(greetHandler))
	mux.Handle("/", simpleMiddleware)

	server := &http.Server{
		Addr:    ":7166",
		Handler: mux,
	}

	fmt.Println(server.ListenAndServe())
}
