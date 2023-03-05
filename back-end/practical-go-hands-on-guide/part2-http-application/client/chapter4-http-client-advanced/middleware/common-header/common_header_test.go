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
				// 요청 헤더를 그대로 응답 헤더로 설정
				for key, value := range r.Header {
					w.Header().Set(key, value[0])
				}

				fmt.Fprint(w, "Just echo the request header")
			}))

	return ts
}

func TestCommonHeaderMiddleware(t *testing.T) {
	testHeader := map[string]string{
		"X-Client-Id": "yjchoi7166",
		"X-Auth-Hash": "adfjllkajfld",
		"X-Token":     "token",
	}

	client := CreateClient(testHeader)

	mock := createMockHttpServer()
	defer mock.Close()

	res, err := client.Get(mock.URL)
	if err != nil {
		t.Fatal("Expected error to be non-nil, got nil")
	}

	for key, value := range testHeader {
		if res.Header.Get(key) != testHeader[key] {
			t.Fatalf("Expected header: %s:%s, Got: %s:%s", key, value, key, testHeader[key])
		}
	}

	t.Logf("\nRequest Header:\n%#v\nResponse Header:\n%#v\n", testHeader, res.Header)
}
