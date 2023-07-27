package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

func main() {
	// 구글에서 전송한 인증서를 운영체제 내 인증 저장소를 이용하여 신뢰성 검증
	conn, err := tls.DialWithDialer(
		&net.Dialer{Timeout: 30 * time.Second},
		"tcp",
		"www.google.com:443",
		&tls.Config{
			CurvePreferences: []tls.CurveID{tls.CurveP256},
			MinVersion:       tls.VersionTLS12,
		},
	)
	if err != nil {
		panic(err)
	}
	defer func() { _ = conn.Close() }()

	state := conn.ConnectionState()
	fmt.Printf("TLS Version: 1.%d\n", state.Version-tls.VersionTLS10)
	fmt.Printf("Cipher Suite Name: %s\n", tls.CipherSuiteName(state.CipherSuite))
	fmt.Printf("TLS Issuer: %s\n", state.VerifiedChains[0][0].Issuer.Organization[0])
}
