package client

import (
	"crypto/tls"
	"net/http"
)

// NopTransport 아무 일도 하지 않는 RoundTripper 구현체
// RoundTripper: 클라이언트와 목적지서버까지의 요청 및 응답 정보를 포함하는 인터페이스
// => TCP 연결 관리
type NopTransport struct{}

func (n *NopTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: http.StatusTeapot}, nil
}

func SetDefaultClient(isSecure, noOperation bool) *http.Client {
	client := http.DefaultClient

	// turn of the SSL
	if !isSecure {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		}
	}

	if noOperation {
		client.Transport = &NopTransport{}
	}

	http.DefaultClient = client
	return client
}
