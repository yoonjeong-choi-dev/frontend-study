package main

import (
	"bytes"
	"fmt"
	"log"
)

func LogExample() {
	buf := bytes.Buffer{}

	logger := log.New(&buf, "[logger] ", log.Lshortfile|log.Ldate)
	logger.Print("test")

	logger.SetPrefix("[new logger] ")
	logger.Printf("you can also add args(%v) and use log.Fatalln to log and crash", true)

	fmt.Printf("Log Result:\n%s\n", buf.String())
}
