package main

import (
	"io"
	"log"
	"net"
)

func main() {
	// serverListener for server
	serverListener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("Listening with %s\n", serverListener.Addr().String())

	// Server side listening
	done := make(chan struct{})
	go func() {
		defer func() {
			done <- struct{}{}
		}()

		for {
			conn, err := serverListener.Accept()
			if err != nil {
				log.Printf("FIN from Client -> Terminate: %s\n", err.Error())
				return
			}

			// Go Handler - Business Logic
			// : Log 1024 bytes received data
			go func(c net.Conn) {
				// graceful termination for TCP Connection
				defer func() {
					// Send FIN to client
					_ = c.Close()

					// process only one handler
					// 계속 리스닝하려면 아래 코드 주석
					done <- struct{}{}
				}()

				// read data from client
				buf := make([]byte, 1024)
				for {
					size, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							log.Fatalf("Error for reading tcp data: %s\n", err.Error())
						}

						// io.EOF : FIN from Client
						return
					}
					log.Printf("Received Data :%q\n", buf[:size])
				}
			}(conn)
		}
	}()

	clientConn, err := net.Dial("tcp", serverListener.Addr().String())
	if err != nil {
		log.Fatalln(err)
	}

	sendData := []byte("Test Data")
	_, err = clientConn.Write(sendData)
	if err != nil {
		log.Printf("Error for recieving data: %s\n", err.Error())
	}

	_ = clientConn.Close()
	<-done

	_ = serverListener.Close()
	<-done
}
