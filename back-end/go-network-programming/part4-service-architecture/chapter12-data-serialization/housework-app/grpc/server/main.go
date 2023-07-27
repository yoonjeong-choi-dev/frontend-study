package main

import (
	"crypto/tls"
	"fmt"
	"google.golang.org/grpc"
	"housework-app/housework/v1"
	"log"
	"net"
)

var (
	addr     = "localhost:7166"
	certFile = "./certs/cert.pem"
	keyFile  = "./certs/key.pem"
)

func main() {
	server := grpc.NewServer()
	robotMaid := new(RobotMaid)
	housework.RegisterRobotMaidServer(server, robotMaid)

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listening for TLS connnection on %s\n", addr)
	log.Fatalln(server.Serve(
		tls.NewListener(
			listener,
			&tls.Config{
				Certificates:             []tls.Certificate{cert},
				CurvePreferences:         []tls.CurveID{tls.CurveP256},
				MinVersion:               tls.VersionTLS12,
				PreferServerCipherSuites: true,
			})))
}
