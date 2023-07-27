package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	addr := "localhost:7166"
	maxIdle := time.Second
	server := NewTLSServer(ctx, addr, maxIdle, nil)

	done := make(chan struct{})
	go func() {
		err := server.ListenAndServeTLS("./certs/cert.pem", "./certs/key.pem")
		if err != nil {
			log.Fatalln(err)
			return
		}
		done <- struct{}{}
	}()

	server.Ready()
	<-done
}
