package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type LoggingMiddleware struct {
	log *log.Logger
}

// RoundTrip : Implement the interface RoundTripper
func (middleware *LoggingMiddleware) RoundTrip(req *http.Request) (*http.Response, error) {
	middleware.log.Printf(
		"Sending a %s request to %s over %s\n",
		req.Method, req.URL, req.Proto,
	)

	res, err := http.DefaultTransport.RoundTrip(req)
	middleware.log.Printf("Got back a response over %s\n", res.Proto)

	return res, err
}

func fetchRemoteResource(client *http.Client, url string) ([]byte, error) {
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "Must specify a HTTP URL to fetch")
		os.Exit(1)
	}

	customTransport := LoggingMiddleware{log: log.New(os.Stdout, "", log.LstdFlags)}

	// create a client with middleware
	client := &http.Client{Timeout: 5 * time.Second, Transport: &customTransport}

	body, err := fetchRemoteResource(client, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "Bytes in response: %d\n", len(body))
}
