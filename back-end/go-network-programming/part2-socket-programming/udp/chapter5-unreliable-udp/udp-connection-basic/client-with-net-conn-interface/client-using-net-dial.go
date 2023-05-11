package main

import (
	"context"
	"fmt"
	"net"
	"server"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// udp 통신을 위한 리스너(서버)는 반드시 net.PacketConn 객체를 이용해야 함
	serverAddr, err := server.EchoUDPServer(ctx, "127.0.0.1:")
	if err != nil {
		panic(err)
	}
	defer cancel()

	// net.Dial 함수를 이용하여 net.Conn 객체 이용
	client, err := net.Dial("udp", serverAddr.String())
	if err != nil {
		panic(err)
	}
	defer func() { _ = client.Close() }()

	// interrupt test cf) pinning-cert-client-with-multi-servers
	interloper, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		panic(err)
	}

	interruptData := []byte("Can you see this...?")
	size, err := interloper.WriteTo(interruptData, client.LocalAddr())
	if err != nil {
		panic(err)
	}
	_ = interloper.Close()

	if l := len(interruptData); l != size {
		fmt.Printf("wrote %d bytes of %d\n", size, l)
	}

	// Client -> Echo Server
	// net.Conn 객체이므로, 서버에 대한 주소가 필요 없어짐(생성 할 때 넘겨줌)
	req := []byte("From Client to Echo Server")
	_, err = client.Write(req)
	if err != nil {
		panic(err)
	}

	// Response : interloper 메시지는 받지 않고 에코 서버 응답만 받음
	buf := make([]byte, 1024)
	size, err = client.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response Data: %s\n", string(buf[:size]))

	// Wait for response from interloper
	err = client.SetDeadline(time.Now().Add(time.Second))
	if err != nil {
		panic(err)
	}

	_, err = client.Read(buf)
	fmt.Printf("Response from interloper: %s\n", err.Error())

}
