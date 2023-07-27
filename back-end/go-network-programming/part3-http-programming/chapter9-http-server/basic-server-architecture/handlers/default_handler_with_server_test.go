package handlers

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestDefaultMethodsHandlerByHTTPServer(t *testing.T) {
	// 서버 초기화
	server := &http.Server{
		Addr: "127.0.0.1:7166",
		// TimoutHandler: 핸들러의 처리 시간 컨트롤
		Handler: http.TimeoutHandler(
			DefaultHandler(), 2*time.Minute, "timeout...."),
		// IdleTimeout: TCP Keep-Alive 설정인 경우, 다음 요청을 대기하기 위해 소켓을 열어두는 시간
		IdleTimeout: 5 * time.Minute,
		// ReadHeaderTimeout: 서버가 요청 헤더를 읽기 위한 최대 시간
		ReadHeaderTimeout: time.Minute,
	}

	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		t.Fatal(err)
	}

	// tcp listener 객체를 이용하여 http server open
	go func() {
		err := server.Serve(listener)
		// 정상 종료가 아닌 경우 로깅
		if err != http.ErrServerClosed {
			t.Fatal(err)
		}
	}()

	// Test Cases
	testCases := []struct {
		title    string
		method   string
		body     io.Reader
		code     int
		response string
	}{
		{"Get", http.MethodGet, nil, http.StatusOK, "Hello, anonymous!"},
		{"Post", http.MethodPost, bytes.NewBufferString("<Script Test>"),
			http.StatusOK, "Hello, &lt;Script Test&gt;!"},
		{"Options", http.MethodOptions, nil,
			http.StatusMethodNotAllowed, "Method not allowed"},
		{"Head", http.MethodHead, nil,
			http.StatusMethodNotAllowed, ""},
	}

	// Client for test
	client := new(http.Client)
	host := fmt.Sprintf("http://%s/", server.Addr)

	for _, test := range testCases {
		req, err := http.NewRequest(test.method, host, test.body)
		if err != nil {
			t.Errorf("[%s] Fail to create request: %v\n", test.title, err)
			continue
		}

		res, err := client.Do(req)
		if err != nil {
			t.Errorf("[%s] Fail to request: %v\n", test.title, err)
		}
		if res.StatusCode != test.code {
			t.Errorf("[%s] expected: %d, got: %d\n", test.title, test.code, res.StatusCode)
		}
		t.Logf("[%s] Response status: %s\n", test.title, res.Status)

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("[%s] Fail to read response body: %v\n", test.title, err)
			continue
		}
		_ = res.Body.Close()

		resBody := strings.TrimSpace(string(data))
		if resBody != test.response {
			t.Errorf("[%s] expected: %s, got: %s\n", test.title, test.response, resBody)
		}
		t.Logf("[%s] Response Body: %s\n", test.title, resBody)
	}

	// close server
	if err := server.Close(); err != nil {
		t.Fatalf("fail to close server: %v\n", err)
	}
}
