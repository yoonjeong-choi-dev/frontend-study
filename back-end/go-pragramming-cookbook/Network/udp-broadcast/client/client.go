package main

import (
	"fmt"
	"net"
)

const addr = "localhost:7166"

func main() {
	addr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer func() { _ = conn.Close() }()
	fmt.Printf("client for server url: %s\n", addr)

	msg := make([]byte, 512)
	_, err = conn.Write([]byte("test-request"))
	if err != nil {
		panic(err)
	}

	for {
		size, err := conn.Read(msg)
		if err == nil {
			fmt.Printf("Response: %s\n", string(msg[:size]))
		}
	}
}
