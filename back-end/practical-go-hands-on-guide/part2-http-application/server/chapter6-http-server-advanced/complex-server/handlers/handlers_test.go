package handlers

import (
	"bytes"
	"complex-server/config"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiHandler(t *testing.T) {
	// mock server io.Writer
	buffer := new(bytes.Buffer)
	conf := config.InitConfig(buffer)

	// mock (request, response)
	req := httptest.NewRequest("GET", "/api", nil)
	resWriter := httptest.NewRecorder()

	apiHandler(resWriter, req, conf)

	res := resWriter.Result()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v\n", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected response status: %v\nGot: %v\n", http.StatusOK, res.StatusCode)
	}

	bodyStr := string(body)
	expectedBody := "This is API handler!"
	if bodyStr != expectedBody {
		t.Errorf("Expected response: %s\nGot: %s\n", expectedBody, bodyStr)
	}
}

// Exercise 6.3
func TestHealthCheckHandler(t *testing.T) {
	tests := []struct {
		method string
		status int
		body   string
	}{
		{
			method: "GET",
			status: http.StatusOK,
			body:   "OK",
		},
		{
			method: "PUT",
			status: http.StatusMethodNotAllowed,
			body:   "Method not allowed\n",
		},
		{
			method: "POST",
			status: http.StatusMethodNotAllowed,
			body:   "Method not allowed\n",
		},
	}

	for _, tc := range tests {
		// mock server io.Writer
		buffer := new(bytes.Buffer)
		conf := config.InitConfig(buffer)

		// mock (request, response)
		req := httptest.NewRequest(tc.method, "/health", nil)
		resWriter := httptest.NewRecorder()

		healthCheckHandler(resWriter, req, conf)

		res := resWriter.Result()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Error reading response body: %v\n", err)
		}

		if res.StatusCode != tc.status {
			t.Errorf("Expected response status: %v\nGot: %v\n", tc.status, res.StatusCode)
		}

		bodyStr := string(body)
		if bodyStr != tc.body {
			t.Errorf("Expected response: %s\nGot: %s\n", tc.body, bodyStr)
		}
	}
}
