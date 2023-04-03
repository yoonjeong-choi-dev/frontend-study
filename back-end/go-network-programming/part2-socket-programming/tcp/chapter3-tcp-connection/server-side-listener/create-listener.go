package main

import (
	"log"
	"net"
)

func main() {
	// :0 => assigned to random port
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatalf("Error for create listener: %s\n", err.Error())
	}

	// graceful termination
	defer func() {
		_ = listener.Close()
	}()

	log.Printf("bound to %q\n", listener.Addr())
}
