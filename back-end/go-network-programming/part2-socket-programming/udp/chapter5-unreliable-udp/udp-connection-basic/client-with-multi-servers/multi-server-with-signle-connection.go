package main

import (
	"context"
	"fmt"
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

	// 클라이언트와 통신할 두 번째 서버
	// => 클라이언트의 udp 연결 객체에 데이터 송신
	interloper, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		panic(err)
	}

	interruptData := []byte("Interrupt...!")
	size, err := interloper.WriteTo(interruptData, client.LocalAddr())
	if err != nil {
		panic(err)
	}
	_ = interloper.Close()

	if l := len(interruptData); l != size {
		fmt.Printf("wrote %d bytes of %d\n", size, l)
	}

	// 클라이언트 -> 에코 서버
	req := []byte("message with echo server")
	_, err = client.WriteTo(req, serverAddr)
	if err != nil {
		panic(err)
	}

	// Response 1 from interloper
	// : TCP 통신의 경우 두 호스트 간 세션이 존재하여 해당 응답을 받을 수 없음
	buf := make([]byte, 1024)
	size, resAddr, err := client.ReadFrom(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response from Interloper")
	fmt.Printf("Server Addr: %s, Response Addr: %s\n", interloper.LocalAddr().String(), resAddr.String())
	fmt.Printf("Response: %s\n", string(buf[:size]))

	// Response 2 from echo server
	size, resAddr, err = client.ReadFrom(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nResponse from Echo Server")
	fmt.Printf("Server Addr: %s, Response Addr: %s\n", serverAddr.String(), resAddr.String())
	fmt.Printf("Response: %s\n", string(buf[:size]))

}
