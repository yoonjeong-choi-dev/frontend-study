package main

import (
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
)

const addr = "ws://localhost:7166/"

// catchSig graceful termination of websocket
func catchSig(ch chan os.Signal, conn *websocket.Conn) {
	<-ch
	err := conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("failed to write close conn:", err)
	}
	return
}

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	log.Printf("connecting to %s\n", addr)

	conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		log.Fatalln("failed to connect:", err)
	}
	defer func() { _ = conn.Close() }()

	go catchSig(interrupt, conn)

	process(conn)
}
