package main

import (
	"fmt"
	"metrics"
	"net/http"
)

func main() {
	http.HandleFunc("/counter", metrics.CounterHandler)
	http.HandleFunc("/timer", metrics.TimerHandler)
	http.HandleFunc("/report", metrics.ReportHandler)

	fmt.Println(http.ListenAndServe(":7166", nil))
}
