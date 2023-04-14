package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const addr = "localhost:7166"

func main() {
	// 사용자 입력을 받기 위한 리더
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter some text('quit' to exit): ")
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error reading input: %s\n", err.Error())
			continue
		}

		if data == "quit\n" {
			fmt.Println("Quit...")
			break
		}

		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Printf("error creating connection: %s\n", err.Error())
			break
		}

		// send data
		fmt.Fprintf(conn, data)

		// get response
		res, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("error reading response: %s\n", err.Error())
		} else {
			fmt.Printf("Response: %s\n", res)
		}
		conn.Close()
	}
}
