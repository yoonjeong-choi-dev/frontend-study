package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		response string
	}{
		{
			name:     "index",
			path:     "/api",
			response: "Hello, this is a simple server",
		},
		{
			name:     "health check",
			path:     "/health",
			response: "OK",
		},
	}

	// Init server
	mux := http.NewServeMux()
	registerHandler(mux)

	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res, err := http.Get(testServer.URL + tc.path)
			defer res.Body.Close()

			if err != nil {
				log.Fatal(err)
			}

			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}

			bodyStr := string(resBody)
			if bodyStr != tc.response {
				t.Errorf("\nExpected: %s\nGot: %s\n", tc.response, bodyStr)
			}
		})
	}

}
