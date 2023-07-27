package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"metric-code-instrumentation/metrics"
	"net"
	"net/http"
	"sync"
	"time"
)

const (
	metricsAddr = "127.0.0.1:8081"
	webAddr     = "127.0.0.1:7166"
)

// newHTTPServer 메트릭 서버 및 웹 서버를 동시에 생성하기 위한 유틸 함수
func newHTTPServer(wg *sync.WaitGroup, addr string, mux http.Handler, stateFunc func(net.Conn, http.ConnState)) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		IdleTimeout:       time.Minute,
		ReadHeaderTimeout: 30 * time.Second,
		ConnState:         stateFunc,
	}

	go func() {
		defer func() {
			wg.Done()
		}()
		wg.Add(1)
		fmt.Println(server.Serve(l))
	}()

	return nil
}

func goodHandler(w http.ResponseWriter, _ *http.Request) {
	metrics.Requests.Add(1)
	defer func(start time.Time) {
		duration := time.Since(start).Seconds()
		metrics.RequestDuration.Observe(duration)
		metrics.RequestDurationSummary.Observe(duration)
	}(time.Now())

	_, err := w.Write([]byte("Good!"))
	if err != nil {
		metrics.WriteErrors.Add(1)
	}
}

func badHandler(w http.ResponseWriter, _ *http.Request) {
	metrics.Requests.Add(1)
	defer func(start time.Time) {
		duration := time.Since(start).Seconds()
		metrics.RequestDuration.Observe(duration)
		metrics.RequestDurationSummary.Observe(duration)
	}(time.Now())

	// always internal server error
	metrics.WriteErrors.Add(1)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte("Bad"))
}

func connStateMetrics(_ net.Conn, state http.ConnState) {
	switch state {
	case http.StateNew:
		metrics.OpenConnections.Add(1)
	case http.StateClosed:
		metrics.OpenConnections.Add(-1)
	}
}

func main() {
	wg := new(sync.WaitGroup)

	metricMux := http.NewServeMux()
	metricMux.Handle("/metrics", promhttp.Handler())
	if err := newHTTPServer(wg, metricsAddr, metricMux, nil); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Metric server listening on %s...\n", metricsAddr)

	webMux := http.NewServeMux()
	webMux.HandleFunc("/good", goodHandler)
	webMux.HandleFunc("/bad", badHandler)
	if err := newHTTPServer(wg, webAddr, webMux, connStateMetrics); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Metric server listening on %s...\n", webAddr)

	wg.Wait()
}
