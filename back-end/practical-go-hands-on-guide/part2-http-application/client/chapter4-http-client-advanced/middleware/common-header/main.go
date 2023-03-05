package main

import (
	"net/http"
)

type CommonHeaderMiddleware struct {
	headers map[string]string
}

// RoundTrip : Implement the interface RoundTripper
func (middleware *CommonHeaderMiddleware) RoundTrip(req *http.Request) (*http.Response, error) {
	requestCopy := req.Clone(req.Context())

	for key, value := range middleware.headers {
		requestCopy.Header.Set(key, value)
	}

	return http.DefaultTransport.RoundTrip(requestCopy)
}

func CreateClient(headers map[string]string) *http.Client {
	middleware := CommonHeaderMiddleware{
		headers: headers,
	}

	client := http.Client{
		Transport: &middleware,
	}

	return &client
}
