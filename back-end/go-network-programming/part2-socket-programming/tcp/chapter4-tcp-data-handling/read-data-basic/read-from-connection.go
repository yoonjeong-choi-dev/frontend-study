package main

import (
	"crypto/rand"
	"io"
	"log"
	"net"
)

func main() {
	// generate random data with size 16MB
	payload := make([]byte, 1<<24)
	_, err := rand.Read(payload)
	if err != nil {
		panic(err)
	}

	server, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		panic(err)
	}

	// server listening
	go func() {
		conn, err := server.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		defer func() {
			_ = conn.Close()
		}()

		// Write random data for response
		_, err = conn.Write(payload)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	// Main: read data from server
	client, err := net.Dial("tcp", server.Addr().String())
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = client.Close()
	}()

	// make buffer to read response with size 512KB
	step := 1
	buf := make([]byte, 1<<19)
	for {
		size, err := client.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		log.Printf("[Step %d]Read %d bytes\n", step, size)
		step++
	}
}
