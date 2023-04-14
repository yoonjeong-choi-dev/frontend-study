package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type connections struct {
	clients map[string]*net.UDPAddr
	mu      sync.Mutex // 고루틴 핸들러에서 동시에 작업하므로 뮤텍스 사용
}

func broadcast(conn *net.UDPConn, c *connections) {
	count := 0
	for {
		count++
		c.mu.Lock()

		for _, client := range c.clients {
			// message for broadcast
			msg := fmt.Sprintf("Sent %d", count)
			if _, err := conn.WriteToUDP([]byte(msg), client); err != nil {
				fmt.Printf("error broadcasting: %s\n", err.Error())
				continue
			}
		}

		c.mu.Unlock()
		time.Sleep(1 * time.Second)
	}
}
