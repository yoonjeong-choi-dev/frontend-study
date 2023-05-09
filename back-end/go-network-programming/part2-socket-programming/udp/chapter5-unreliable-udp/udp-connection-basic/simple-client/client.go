package main

import (
	"context"
	"log"
	"net"
	"server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	serverAddr, err := server.EchoUDPServer(ctx, "127.0.0.1:")
	if err != nil {
		panic(err)
	}
	defer cancel()

	client, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		panic(err)
	}
	defer func() { _ = client.Close() }()

	msg := []byte("Simple Client")
	_, err = client.WriteTo(msg, serverAddr)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)
	size, resAddr, err := client.ReadFrom(buf)
	if err != nil {
		panic(err)
	}

	// 세션 개념이 없으므로, 요청한 주소와 응답한 주소에 대한 검증 필요
	log.Printf("Server Addr: %s, Response Addr: %s\n", serverAddr.String(), resAddr.String())
	log.Printf("Response: %s\n", string(buf[:size]))
}
