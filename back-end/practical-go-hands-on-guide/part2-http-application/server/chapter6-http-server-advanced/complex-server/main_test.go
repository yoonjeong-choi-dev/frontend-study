package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSetupServer(t *testing.T) {
	buffer := new(bytes.Buffer)
	mux := http.NewServeMux()
	wrappedMux := setupServer(mux, buffer)

	testServer := httptest.NewServer(wrappedMux)
	defer testServer.Close()

	res, err := http.Get(testServer.URL + "/panic")
	if err != nil {
		t.Errorf("GET panic - error: %v\n", err)
	}
	defer res.Body.Close()

	// body 쪽은 middleware 에서 테스트
	_, err = io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v\n", err)
	}

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected response status: %v\nGot: %v\n", http.StatusInternalServerError, res.StatusCode)
	}

	// panicHandlingMiddleware 가 남긴 로그 확인
	logs := buffer.String()
	expectedLogs := []string{
		"protocol:HTTP/1.1, path:/panic, method:GET, duration:",
		"panic detected",
	}

	for _, log := range expectedLogs {
		if !strings.Contains(logs, log) {
			t.Errorf(
				"Expected logs to contain: %s\nGot: %s\n",
				log, logs,
			)
		}
	}
}
