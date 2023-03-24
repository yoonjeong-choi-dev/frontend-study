package main

import (
	"bytes"
	"fmt"
	io_reader_writer "io-reader-writer"
	"log"
)

func main() {
	fmt.Println("Basic I/O Interface Usage\n")

	fmt.Println("Example 1: Copy from reader to writer interface")
	testData := "Let's copy~~"
	in := bytes.NewReader([]byte(testData))
	out := &bytes.Buffer{}

	fmt.Printf("in(io.Reader) : %s\n", testData)
	fmt.Print("os.stdout on Copy: ")
	if err := io_reader_writer.Copy(in, out); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	fmt.Printf("out(io.Writer) output: %s\n", out.String())

	fmt.Println("\nExample2: io.Pipe()")
	fmt.Print("os.stdout on Pipe: ")
	if err := io_reader_writer.PipeExample(); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
