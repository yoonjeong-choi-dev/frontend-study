package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const addr = "localhost:7166"

func echoHandler(conn net.Conn) {
	req := bufio.NewReader(conn)
	data, err := req.ReadString('\n')
	if err != nil {
		fmt.Printf("error reading data: %s\n", err.Error())
		return
	}

	fmt.Printf("Request: %s\n", data)

	res := []byte(strings.ToUpper(data))
	conn.Write(res)
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer func() { _ = listener.Close() }()

	fmt.Printf("listening on %s\n", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error accepting connection: %s\n", err.Error())
			continue
		}

		go echoHandler(conn)
	}
}
