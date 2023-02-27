package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createMockHttpServer() *httptest.Server {
	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello World")
			}))

	return ts
}

func TestFetchRemoteResource(t *testing.T) {
	mock := createMockHttpServer()
	defer mock.Close()

	expected := "Hello World"
	data, err := fetchRemoteResource(mock.URL)

	if err != nil {
		t.Fatal(err)
	}

	if expected != string(data) {
		t.Errorf("Expeteced response to be: %s, Got: %s\n", expected, data)
	}
}
