package restapi

import "net/http"

// APITransport 인증정보가 담긴 RoundTripper 구현체
// TCP 연결 관리
type APITransport struct {
	*http.Transport
	username string
	password string
}

func (t *APITransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// 요청 전송 전에, 인증 수행
	// 여기서 인증 관련 정보(ex 토근)을 갱신할 수 있음
	req.SetBasicAuth(t.username, t.password)
	return t.Transport.RoundTrip(req)
}
