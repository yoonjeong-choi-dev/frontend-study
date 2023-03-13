package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMiddleware(t *testing.T) {
	// mock logging output
	var buf bytes.Buffer

	// test TLS server
	mux := http.NewServeMux()
	logger := log.New(
		&buf, "tls-server",
		log.Lshortfile|log.LstdFlags,
	)
	mainHandler := registerHandlersAndMiddlewares(mux, logger)

	testServer := httptest.NewUnstartedServer(mainHandler)
	testServer.EnableHTTP2 = true
	testServer.StartTLS()

	// TLS client from test server
	client := testServer.Client()
	_, err := client.Get(testServer.URL + "/api")
	if err != nil {
		t.Fatal(err)
	}

	// Check Protocol via logging middleware
	expected := "protocol: HTTP/2.0, path: /api, method: GET"
	logs := buf.String()
	if !strings.Contains(logs, expected) {
		t.Fatalf("\nExpected logs to contain %s\nFound: %s\n", expected, logs)
	}

	t.Logf("\nLog: %s\n", logs)
}
