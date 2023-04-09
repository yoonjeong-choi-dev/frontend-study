package main

import (
	"context"
	"log"
	"net"
	"ping"
	"time"
)

func main() {
	pingServerDone := make(chan struct{})

	pingServer, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		log.Fatalf("Error for create pingServer: %s\n", err.Error())
	}

	pingServerBegin := time.Now()
	go func() {
		defer func() { close(pingServerDone) }()

		conn, err := pingServer.Accept()
		if err != nil {
			log.Printf("Server Shutdown: %s\n", err.Error())
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer func() {
			cancel()
			conn.Close()
		}()

		resetTimer := make(chan time.Duration, 1)
		resetTimer <- time.Second

		// Request Ping by PingWithInterval
		go ping.PingWithInterval(ctx, conn, resetTimer)

		// Set deadline for reading request & writing response
		err = conn.SetDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			log.Printf("Error for setting deadline : %s\n", err.Error())
			return
		}

		buf := make([]byte, 1024)
		for {
			size, err := conn.Read(buf)
			if err != nil {
				log.Printf("Error for reading request: %s\n", err.Error())
				return
			}
			log.Printf("[Ping Server] %s - %s\n",
				buf[:size],
				time.Since(pingServerBegin).Truncate(time.Second),
			)

			// 요청을 받은 경우, 데드라인 및 ping 초기화
			resetTimer <- 0
			err = conn.SetDeadline(time.Now().Add(5 * time.Second))
			if err != nil {
				log.Printf("Error for setting deadline : %s\n", err.Error())
				return
			}
		}
	}()

	pongServer, err := net.Dial("tcp", pingServer.Addr().String())
	if err != nil {
		log.Fatalf("Error for create pongServer: %s\n", err.Error())
	}
	defer pongServer.Close()

	// Read Ping
	buf := make([]byte, 1024)
	for i := 0; i < 4; i++ {
		size, err := pongServer.Read(buf)
		if err != nil {
			log.Fatalf("Disconnection: %s\n", err.Error())
		}

		log.Printf("[Response] %s - %s\n",
			buf[:size],
			time.Since(pingServerBegin).Truncate(time.Second),
		)
	}

	// Response Pong
	// => Ping 서버는 pong 메시지를 받고, 데드라인 초기화
	_, err = pongServer.Write([]byte("PONG"))
	if err != nil {
		log.Fatalf("Error for writing 'Pong': %s\n", err.Error())
	}

	// Read Ping
	for i := 0; i < 4; i++ {
		size, err := pongServer.Read(buf)
		if err != nil {
			log.Fatalf("Disconnection: %s\n", err.Error())
		}

		log.Printf("[Pong Server] %s - %s\n",
			buf[:size],
			time.Since(pingServerBegin).Truncate(time.Second),
		)
	}

	// wait for pingServer done
	<-pingServerDone

	pingServerDuration := time.Since(pingServerBegin)
	log.Printf("Ping Server lives.. %s\n", pingServerDuration)

}
