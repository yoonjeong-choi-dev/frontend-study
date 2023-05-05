package main

import (
	"bufio"
	"log"
	"net"
)

const payload = "THe bigger the interface, the weaker the abstraction, This is test payload!"

func main() {
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

		_, err = conn.Write([]byte(payload))
		if err != nil {
			log.Println(err)
			return
		}
	}()

	client, err := net.Dial("tcp", server.Addr().String())
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = client.Close()
	}()

	// scanner for reading response
	scanner := bufio.NewScanner(client)

	// split the word by space
	scanner.Split(bufio.ScanWords)

	// read data by scanner
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	log.Printf("Splitted Res Data: %#v\n", words)
}
