package main

import (
	"built-in-logger/writer"
	"bytes"
	"fmt"
	"log"
	"os"
)

func MultiWriterLog() {
	inMemory := new(bytes.Buffer)
	w := writer.NewSustainedMultiWriter(inMemory, os.Stdout, os.Stderr)

	l := log.New(w, "[multi]", log.Lshortfile|log.Lmsgprefix)

	fmt.Println("Stdout & Stderr:")
	l.Println("This message is logged to multi resources")

	fmt.Println("\nIn-memory Log:")
	fmt.Println(inMemory.String())
}
