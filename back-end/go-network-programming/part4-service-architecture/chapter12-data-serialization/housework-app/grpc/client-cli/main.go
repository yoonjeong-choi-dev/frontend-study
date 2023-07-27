package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"housework-app/housework/v1"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var addr, certFile string

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			`Usage: %s [add chore, ...|complete #]
    add         add comma-separated chores
    complete    complete designated chore
    list        list all chores

Flags:
`, filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}

	flag.StringVar(&addr, "address", "localhost:7166",
		"server address")
	flag.StringVar(&certFile, "ca-cert", "../../certs/cert.pem", "CA certificate")
}

func main() {
	flag.Parse()

	cert, err := ioutil.ReadFile(certFile)
	if err != nil {
		fmt.Println(err)
	}

	// 인증서 고정
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(cert); !ok {
		fmt.Println("failed to add cert to pool")
		os.Exit(1)
	}

	// 네트워크 연결 생성
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(
			credentials.NewTLS(&tls.Config{
				CurvePreferences: []tls.CurveID{tls.CurveP256},
				MinVersion:       tls.VersionTLS12,
				RootCAs:          certPool,
			})),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client := MaidService{client: housework.NewRobotMaidClient(conn)}
	ctx := context.Background()
	switch strings.ToLower(flag.Arg(0)) {
	case "list":
		err = client.list(ctx)
	case "add":
		err = client.add(ctx, strings.Join(flag.Args()[1:], " "))
	case "complete":
		err = client.complete(ctx, flag.Arg(1))
	default:
		err = errors.New("unsupported command")
	}

	if err != nil {
		fmt.Printf("failed to process: %s\n", err.Error())
		os.Exit(1)
	}

	if flag.Arg(0) != "list" {
		err = client.list(ctx)
		if err != nil {
			fmt.Printf("failed to list: %s\n", err.Error())
			os.Exit(1)
		}
	}
}
