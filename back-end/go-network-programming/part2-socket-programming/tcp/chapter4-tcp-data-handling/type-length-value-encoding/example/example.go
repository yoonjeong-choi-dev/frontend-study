package main

import (
	"fmt"
	"log"
	"net"
	"runtime/debug"
	"type-length-value-encoding/decoder"
	"type-length-value-encoding/encoding"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	// Set Payloads
	p1 := encoding.Binary("BinaryType Test Data")
	p2 := encoding.Binary("TLV Encoding Test")
	p3 := encoding.String("StringType Test Data -> Go Network Programming")
	payloads := []encoding.Payload{&p1, &p2, &p3}

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

		// send all payloads
		for _, p := range payloads {
			size, err := p.WriteTo(conn)
			if err != nil {
				log.Printf("Error for write payload with size %d: %s\n", size, err.Error())
				break
			}
		}
	}()

	client, err := net.Dial("tcp", server.Addr().String())
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = client.Close()
	}()

	// Read 3 payloads
	for i := 0; i < len(payloads); i++ {
		data, err := decoder.Decoder(client)
		if err != nil {
			panic(err)
		}

		log.Printf("[%d payload] %s\n", i, data.String())
	}

}
