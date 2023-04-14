package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("failed to upgrade connection:", err)
		return
	}

	for {
		messageType, payload, err := conn.ReadMessage()
		if err != nil {
			log.Println("failed to read message:", err)
			return
		}

		reqMsg := string(payload)
		log.Printf("received from clinet: %#v\n", reqMsg)
		resMsg := fmt.Sprintf("[Echo] %s", reqMsg)

		if err := conn.WriteMessage(messageType, []byte(resMsg)); err != nil {
			log.Println("failed to write message:", err)
			return
		}
	}
}
