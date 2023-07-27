package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	// 인증서 고정: 서버에서 발급한 인증서를 이용하여 통신
	// => 서버가 발급한 인증서 파일 읽기
	cert, err := ioutil.ReadFile("./certs/cert.pem")
	if err != nil {
		panic(err)
	}

	// 운영체제가 신뢰하는 인증서 저장소에 해당 인증서를 등록
	// => 등록된 인증서를 사용하는 서버와의 연결을 신뢰
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(cert); !ok {
		panic(errors.New("failed to append certificate to pool"))
	}

	// 인증서를 등록한 인증서 풀을 이용하여 TLS 연결 정의
	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.CurveP256},
		MinVersion:       tls.VersionTLS12,
		RootCAs:          certPool,
	}

	client, err := tls.Dial("tcp", "localhost:7166", tlsConfig)
	if err != nil {
		panic(err)
	}
	defer func() { _ = client.Close() }()

	testReq := []string{"test", "hello~"}
	for _, req := range testReq {
		_, err = client.Write([]byte(req))
		if err != nil {
			panic(err)
		}

		buf := make([]byte, 1024)
		n, err := client.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Response of '%s': %s\n", req, string(buf[:n]))
	}
}
