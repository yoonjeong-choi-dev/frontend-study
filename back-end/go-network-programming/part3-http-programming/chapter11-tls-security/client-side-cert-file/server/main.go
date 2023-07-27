package main

import (
	"client-side-cert-file/utils"
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"log"
	"net"
	"strings"
)

var (
	clientCertFile = "./cert/clientCert.pem"
	serverCertFile = "./cert/serverCert.pem"
	serverKeyFile  = "./cert/serverKey.pem"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 클라이언트 인증서를 인증서 풀에 등록
	serverPool, err := utils.CACertPool(clientCertFile)
	if err != nil {
		panic(err)
	}

	// 서버 인증서 및 개인키 등록
	cert, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
	if err != nil {
		panic(err)
	}

	serverConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		// GetConfigForClient: 클라이언트의 IP, Host 이름을 알기 위한 메서드를 정의하는 필드
		// (why?) 서버 측은 클라이언트 인증서를 이용하여 IP, Host 값을 알 수 없음
		// => 인증서 검증 과정에서 클라이언트의 정보를 얻기 위한 방법
		// *tls.ClientHelloInfo 매개변수는 TLS 핸드셰이크 과정에서 생성되는 포인터
		GetConfigForClient: func(hello *tls.ClientHelloInfo) (*tls.Config, error) {
			return &tls.Config{
				Certificates: []tls.Certificate{cert},
				// TLS 핸드셰이크 과정이 끝나기 전에 클라이언트가 유효한 인증서를 제공하였는지 확인
				ClientAuth:               tls.RequireAndVerifyClientCert,
				ClientCAs:                serverPool,
				CurvePreferences:         []tls.CurveID{tls.CurveP256},
				MinVersion:               tls.VersionTLS13,
				PreferServerCipherSuites: true,
				// 인증서 검증 절차를 정의하는 필드: 디폴트 인증서 검증 후에 호출되는 함수
				// => 클라이언트의 호스트 이름 검증
				VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
					opts := x509.VerifyOptions{
						KeyUsages: []x509.ExtKeyUsage{
							x509.ExtKeyUsageClientAuth,
						},
						// 서버 풀에 등록되어 있는 클라이언트 인증서는 신뢰
						Roots: serverPool,
					}

					ip := strings.Split(hello.Conn.RemoteAddr().String(), ":")[0]
					hostnames, err := net.LookupAddr(ip)
					if err != nil {
						// 추적이 실패하면 빈 슬라이스가 반환
						log.Printf("[error] PTR lookup: %v\n", err)
					}

					hostnames = append(hostnames, ip)

					// 모든 인증서 체인(리프 인증서 -> .. -> 최상단 인증 기관) 순회
					for _, chain := range verifiedChains {
						opts.Intermediates = x509.NewCertPool()

						// 중간 인증서 풀에 리프 인증서를 제외한 모든 인증서 추가
						for _, cert := range chain[1:] {
							opts.Intermediates.AddCert(cert)
						}

						// 체인에 있는 리프 인증서를 이용하여 클라이언트 호스트 이름 검증
						// => 인증서 체인의 리프 인증서 중에 호스트 이름이 유효한 경우 검증 완료
						for _, hostname := range hostnames {
							opts.DNSName = hostname
							_, err = chain[0].Verify(opts)
							if err == nil {
								return nil
							}
						}
					}
					// 인증서 체인의 모든 리프 인증서가 현재 클라이언트 호스트를 검증 못하는 경우 인증 불가능
					return errors.New("client authentication failed")
				},
			}, nil
		},
	}

	addr := "localhost:7166"
	server := NewTLSEchoServer(ctx, addr, 0, serverConfig)

	done := make(chan struct{})
	go func() {
		err := server.ListenAndServeTLS(serverCertFile, serverKeyFile)
		if err != nil {
			log.Fatalln(err)
			return
		}
		done <- struct{}{}
	}()

	server.Ready()
	<-done
}
