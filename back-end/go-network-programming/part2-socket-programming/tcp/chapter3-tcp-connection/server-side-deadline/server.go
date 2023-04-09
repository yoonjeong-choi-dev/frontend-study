package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		log.Fatalf("Error for lintener: %v\n", err.Error())
	}

	// listen client request
	sync := make(chan struct{})
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Disconnect: %s\n", err.Error())
			return
		}

		defer func() {
			conn.Close()
			close(sync)
		}()

		// set deadline for reading request 1
		err = conn.SetDeadline(time.Now().Add(5 * time.Second))

		// blocking by client request
		// : 클라이언트에서 요청을 보낼 때까지 블록킹
		buf := make([]byte, 1)
		_, err = conn.Read(buf)
		netErr, ok := err.(net.Error)
		if !ok {
			log.Fatalln("Error for converting error -> net.Error")
		}

		log.Println("First Disconnection")
		log.Printf("Timeout : %v\n", netErr.Timeout())
		log.Printf("Temporary: %v\n", netErr.Temporary())
		log.Printf("Error Message: %s\n", netErr.Error())

		// 첫 데드라인 발생 통보
		// => 이후 밑의 클라이언트(conn)은 제대로 된 요청을 보낸다
		sync <- struct{}{}

		// set deadline for reading request 2
		err = conn.SetDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			log.Printf("Error for setting deadline 2: %s\n", err.Error())
			return
		}

		// Second request read
		buf = make([]byte, 1)
		size, err := conn.Read(buf)
		if err != nil {
			netErr, ok := err.(net.Error)
			if !ok {
				log.Fatalln("Error for converting error -> net.Error")
			}

			log.Println("First Disconnection")
			log.Printf("Timeout : %v\n", netErr.Timeout())
			log.Printf("Temporary: %v\n", netErr.Temporary())
			log.Printf("Error Message: %s\n", netErr.Error())
		} else {
			log.Printf("Read Request: %s\n", string(buf[:size]))
		}
	}()

	// Client
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		log.Fatalf("Error for connection: %s\n", err.Error())
	}

	// First Request: 아무것도 하지 않음
	<-sync

	// Second Request
	_, err = conn.Write([]byte("7"))
	if err != nil {
		log.Fatalf("Error for request: %s\n", err.Error())
	}

	// Wait response for the second request
	// => 아무것도 전송하지 않기 때문에 서버 측 Write Deadline 발생
	res := make([]byte, 1)
	_, err = conn.Read(res)
	log.Printf("io.EOF(disconnect by server) :%v\n", err == io.EOF)
	log.Printf("Error Message: %s\n", err.Error())
}
