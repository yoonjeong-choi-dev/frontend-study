package reverse_proxy

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
)

type Proxy struct {
	// Client Proxy 서버는 요청 처리 서버의 클라이언트
	Client *http.Client

	// ServerURL 실제 요청을 처리하는 서버의 url
	ServerURL string
}

// ServeHTTP http.Handler 인터페이스 구현
func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := p.ProcessRequest(r); err != nil {
		log.Printf("error to process client request: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 서버에 요청
	res, err := p.Client.Do(r)
	if err != nil {
		log.Printf("error to process request from grpc-server: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 서버에게서 받은 응답을 그대로 클라이언트에게 전달
	defer func() { _ = res.Body.Close() }()
	CopyResponse(w, res)
}

func (p *Proxy) ProcessRequest(r *http.Request) error {
	// 현재 클라이언트는 브라우저
	// => 요청 처리 서버에서 응답을 받은 이후, 연쇄적인 요청(정적 리소스)이 발생
	// => 리소스 요청들을 받은 프록시 서버는 다시 서버에 해당 요청들을 응답받아와 전달해야 함
	proxyUrlStr := p.ServerURL + r.URL.String()
	proxyURL, err := url.Parse(proxyUrlStr)
	if err != nil {
		return err
	}

	// 리버스 프록시: 서버 대신 요청을 받아 서버 자체를 클라이언트에서 분리
	// => proxyURL 에 대응하는 서버로 요청처리를 위해 요청 객체 변경
	r.URL = proxyURL
	r.Host = proxyURL.Host
	r.RequestURI = ""
	return nil
}

// CopyResponse 서버에게 받은 응답을 클라이언트의 응답에 복사
func CopyResponse(w http.ResponseWriter, res *http.Response) {
	// Copy Body
	var out bytes.Buffer
	_, _ = out.ReadFrom(res.Body)

	// Copy Header
	for key, values := range res.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(res.StatusCode)

	_, _ = w.Write(out.Bytes())
}
