package main

import (
	"log"
	"net"
	"syscall"
	"time"
)

func CreateConnectionWithDialTimeout(
	protocol, address string,
	timeout time.Duration,
) (net.Conn, error) {
	dialer := net.Dialer{
		Control: func(_, address string, _ syscall.RawConn) error {
			return &net.DNSError{
				Err:         "connection timeout",
				Name:        address,
				Server:      "127.0.0.1",
				IsTimeout:   true,
				IsTemporary: true,
			}
		},
		Timeout: timeout,
	}

	return dialer.Dial(protocol, address)
}

func main() {
	conn, err := CreateConnectionWithDialTimeout("tcp", "10.0.0.1:http", 10*time.Second)
	if err == nil {
		conn.Close()
		log.Fatalln("Connection is closed without timeout")
	}

	// 일반적인 에러를 net.Error로 변환하여 세부적인 에러 파악
	netErr, ok := err.(net.Error)
	if !ok {
		log.Fatalln("Cannot not convert to net.Error")
	}

	log.Printf("Timeout : %v\n", netErr.Timeout())
	log.Printf("Temporary: %v\n", netErr.Temporary())
	log.Printf("Error Message: %s\n", netErr.Error())
}
