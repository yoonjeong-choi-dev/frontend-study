package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"rpc/tweak"
)

const PORT = ":7166"

func main() {
	receiver := new(tweak.StringTweaker)
	if err := rpc.Register(receiver); err != nil {
		log.Fatalln("failed to register:", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalln("failed to listen:", err)
	}

	fmt.Printf("listening on %s\n", PORT)
	log.Panic(http.Serve(listener, nil))
}
