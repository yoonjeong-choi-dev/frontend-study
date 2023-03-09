package middleware

import (
	"bytes"
	"complex-server/config"
	"complex-server/handlers"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPanicHandlingMiddleware(t *testing.T) {
	// mock server - config
	buffer := new(bytes.Buffer)
	conf := config.InitConfig(buffer)

	// mock server - handler
	mux := http.NewServeMux()
	handlers.RegisterHandler(mux, conf)

	// mock server - panic middleware
	handler := panicHandlingMiddleware(mux, conf)

	// mock (request, response)
	req := httptest.NewRequest("GET", "/panic", nil)
	resWriter := httptest.NewRecorder()

	// call api
	handler.ServeHTTP(resWriter, req)

	res := resWriter.Result()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v\n", err)
	}

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected response status: %v\nGot: %v\n", http.StatusInternalServerError, res.StatusCode)
	}

	bodyStr := string(body)
	expectedBody := "unexpected server error"
	if bodyStr != expectedBody {
		t.Errorf("Expected response: %s\nGot: %s\n", expectedBody, bodyStr)
	}
}
