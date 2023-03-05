package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptrace"
	"os"
	"time"
)

func createClientWithTimeout(d time.Duration) *http.Client {
	client := http.Client{Timeout: d}
	return &client
}

func createHttpGetRequestWithConnectionPool(
	ctx context.Context, url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	trace := &httptrace.ClientTrace{
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Get Conn from Pool: %+v\n", connInfo)
		},
	}

	ctxTrace := httptrace.WithClientTrace(req.Context(), trace)
	request := req.WithContext(ctxTrace)
	return request, err
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "Must specify a HTTP URL to fetch")
		os.Exit(1)
	}

	timeout := 5 * time.Second
	ctx := context.Background()
	client := createClientWithTimeout(timeout)

	req, err := createHttpGetRequestWithConnectionPool(ctx, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		os.Exit(1)
	}

	fmt.Println("Infinite Request")
	for {
		client.Do(req)
		time.Sleep(1 * time.Second)
		fmt.Println("------------------------------")
	}
}
