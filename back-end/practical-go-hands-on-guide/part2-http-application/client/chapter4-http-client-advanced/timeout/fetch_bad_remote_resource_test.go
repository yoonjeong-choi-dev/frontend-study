package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// createBadHttpServer 문제점
// 테스트가 time.sleep 으로 인해 그전에 테스트 코드가 끝나도 60초 대기해야함
func createBadHttpServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(60 * time.Second)
			fmt.Fprint(w, "Hello World")
		}))

	return ts
}

func createBadHttpServerImproved(shutdownFlag chan struct{}) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// shutdownFlag 채널에 값을 전달하지 않는 이상 여기서 계속 채널을 읽음
			// => 응답이 무한하게 지연
			<-shutdownFlag
			fmt.Fprint(w, "Hello World")
		}))

	return ts
}

func TestFetchBadRemoteResource(t *testing.T) {
	//mock := createBadHttpServer()

	shutdownFlag := make(chan struct{})
	mock := createBadHttpServerImproved(shutdownFlag)
	defer mock.Close()
	defer func() {
		// 함수 종료 시, 서버도 종료 시킴
		shutdownFlag <- struct{}{}
	}()

	client := createClientWithTimeout(200 * time.Millisecond)
	_, err := fetchRemoteResource(client, mock.URL)
	if err == nil {
		t.Fatalf("Expected non-nil error")
	}

	expectedErrMsg := "context deadline exceeded"
	if !strings.Contains(err.Error(), expectedErrMsg) {
		t.Fatalf("\nExpected error to contain: %s, but\nGot: %v\n",
			expectedErrMsg, err)
	}

	t.Logf("\nTimeout Error: %v\n", err)
}
