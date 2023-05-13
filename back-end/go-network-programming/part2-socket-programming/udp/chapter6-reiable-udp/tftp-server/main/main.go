package main

import (
	"io/ioutil"
	"log"
	"os"
	tftp_server "tftp-server"
)

func main() {
	addr := "127.0.0.1:7166"
	fileName := "payload.svg"

	path, _ := os.Getwd()
	log.Println(path)

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}

	server := tftp_server.Server{Payload: file}
	log.Fatal(server.ListenAndServe(addr))
}
