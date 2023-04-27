package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Simple Server is on with :80")
	log.Fatalln(http.ListenAndServe(":80", nil))
}
