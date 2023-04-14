package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"strings"
)

func process(conn *websocket.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter a message: ")
		data, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read stdin:", err)
			return
		}

		data = strings.TrimSpace(data)
		if data == "quit" || data == "q" {
			return
		}

		err = conn.WriteMessage(websocket.TextMessage, []byte(data))
		if err != nil {
			log.Println("failed to write message:", err)
			return
		}

		_, res, err := conn.ReadMessage()
		if err != nil {
			log.Println("failed to read message:", err)
			return
		}

		log.Printf("Recieved from server: %#v\n", string(res))
	}
}
