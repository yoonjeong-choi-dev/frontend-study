package main

import (
	"context"
	"log"
	"net"
	"syscall"
	"time"
)

func main() {
	duration := 5 * time.Second
	deadline := time.Now().Add(duration)

	// create context with deadline
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	var dialer net.Dialer
	dialer.Control = func(_, _ string, _ syscall.RawConn) error {
		// wait to make timeout
		time.Sleep(duration + time.Millisecond)
		return nil
	}

	// Create connection with context
	conn, err := dialer.DialContext(ctx, "tcp", "10.0.0.0:80")
	if err == nil {
		conn.Close()
		log.Fatalln("Connection is closed without timeout")
	}

	// net.Error: error for connection
	netErr, ok := err.(net.Error)
	if !ok {
		log.Fatalln("Cannot not convert to net.Error")
	}

	log.Printf("Timeout : %v\n", netErr.Timeout())
	log.Printf("Temporary: %v\n", netErr.Temporary())
	log.Printf("Error Message: %s\n", netErr.Error())

	// error for context
	log.Printf("DeadlineExceeded: %v\n", ctx.Err() == context.DeadlineExceeded)
	log.Printf("Context Error: %s\n", ctx.Err().Error())
}
