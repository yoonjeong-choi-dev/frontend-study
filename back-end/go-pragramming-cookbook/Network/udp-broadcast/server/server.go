package main

import (
	"fmt"
	"net"
)

const addr = "localhost:7166"

func main() {
	c := &connections{
		clients: make(map[string]*net.UDPAddr),
	}

	addr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer func() { _ = conn.Close() }()
	fmt.Printf("listening on %s\n", addr)

	go broadcast(conn, c)

	msg := make([]byte, 1024)
	for {
		_, clientAddr, err := conn.ReadFromUDP(msg)
		if err != nil {
			fmt.Printf("error reading request: %s\n", err.Error())
			continue
		}

		// 요청 주소 저장
		c.mu.Lock()
		c.clients[clientAddr.String()] = clientAddr
		fmt.Printf("clients list:\n%v\n", c.clients)
		c.mu.Unlock()

		fmt.Printf("%s connected\n", clientAddr)
	}
}
