package main

import (
	"context"
	"log"
	"net"
	"syscall"
	"time"
)

func main() {
	// create context which can disconnect connection by cancel
	ctx, cancel := context.WithCancel(context.Background())
	sync := make(chan struct{})

	// go routine for request
	go func() {
		defer func() {
			sync <- struct{}{}
		}()

		var dialer net.Dialer
		dialer.Control = func(_, _ string, _ syscall.RawConn) error {
			time.Sleep(time.Second)
			return nil
		}

		conn, err := dialer.DialContext(ctx, "tcp", "10.0.0.0:80")
		if err != nil {
			log.Printf("Terminate go routine func: %s\n", err.Error())
			return
		}

		conn.Close()
		log.Println("Connection did not timeout")
	}()

	// disconnect by context cancel
	cancel()

	// by cancel(), stop the go routine
	<-sync

	// error for context
	log.Printf("Canceled: %v\n", ctx.Err() == context.Canceled)
	log.Printf("DeadlineExceeded: %v\n", ctx.Err() == context.DeadlineExceeded)
	log.Printf("Context Error: %s\n", ctx.Err().Error())
}
