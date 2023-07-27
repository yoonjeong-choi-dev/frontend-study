package tls_support_test

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/http2"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func createTLSServer() *httptest.Server {
	// NewTLSServer: 인증서 생성 및 TLS 설정등을 내부적으로 설정
	// => 기본적으로 반환하는 서버의 Client 메서드로 반환되는 클라이언트만 TLS 사용 가능
	return httptest.NewTLSServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.TLS == nil {
				// 요청이 TLS 가 아닌 경우, TLS url 리다이렉션
				redirectPath := fmt.Sprintf("https://%s%s", r.Host, r.RequestURI)
				http.Redirect(w, r, redirectPath, http.StatusMovedPermanently)
				return
			}
			w.WriteHeader(http.StatusOK)
		}))
}

func TestClientFromTestServer(t *testing.T) {
	server := createTLSServer()
	defer server.Close()

	// Test 1. NewTLSServer().Client() 통신
	res, err := server.Client().Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	got := res.StatusCode
	if got != http.StatusOK {
		t.Errorf("expected: %d, got %d\n", http.StatusOK, got)
	}
}

func TestClientWithAuthentication(t *testing.T) {
	server := createTLSServer()
	defer server.Close()

	// TLS 통신을 위한 클라이언트의 Transport 구조체 정의
	// => 암호화 방식으로 타원곡선 사용하여 키 생성
	tp := &http.Transport{
		TLSClientConfig: &tls.Config{
			CurvePreferences: []tls.CurveID{tls.CurveP256},
			MinVersion:       tls.VersionTLS12,
		},
	}

	// 기본 TLS 구성이 아니므로, 명시적으로 HTTP/2 프로토콜 사용하도록 해야 함
	err := http2.ConfigureTransport(tp)
	if err != nil {
		t.Fatal(err)
	}

	// TLS 설정이 완료된 클라이언트 생성
	client := &http.Client{Transport: tp}

	// Test 1: 신뢰할 인증서 설정을 하지 않음
	// => 운영체제가 신뢰하는 인증 저장소의 인증서를 신뢰
	// => 위에서 설정한 Transport 구조체는 서버가 보낸 인증서의 서명자를 신뢰하지 않아 에러 발생
	_, err = client.Get(server.URL)
	if err == nil || !strings.Contains(err.Error(), "certificate is not trusted") {
		t.Errorf("expected: ,got: %q\n", err)
	}

	// Test 2: InsecureSkipVerify 설정
	// => 모든 서버의 인증서를 검증없이 신뢰
	tp.TLSClientConfig.InsecureSkipVerify = true
	res, err := client.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	got := res.StatusCode
	if got != http.StatusOK {
		t.Errorf("expected: %d, got %d\n", http.StatusOK, got)
	}
}
