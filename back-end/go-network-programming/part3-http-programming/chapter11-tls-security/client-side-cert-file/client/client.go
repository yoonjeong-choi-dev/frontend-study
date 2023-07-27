package main

import (
	"client-side-cert-file/utils"
	"crypto/tls"
	"fmt"
)

var (
	clientCertFile = "./cert/clientCert.pem"
	clientKeyFile  = "./cert/clientKey.pem"
	serverCertFile = "./cert/serverCert.pem"
)

func main() {
	// 서버 인증서를 인증서 풀에 등록
	clientPool, err := utils.CACertPool(serverCertFile)
	if err != nil {
		panic(err)
	}

	cert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		panic(err)
	}

	clientConfig := &tls.Config{
		Certificates:     []tls.Certificate{cert},
		CurvePreferences: []tls.CurveID{tls.CurveP256},
		MinVersion:       tls.VersionTLS13,
		RootCAs:          clientPool,
	}

	client, err := tls.Dial("tcp", "localhost:7166", clientConfig)
	if err != nil {
		panic(err)
	}
	defer func() { _ = client.Close() }()

	testReq := []string{"client side cert", "another request with TLS"}
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
