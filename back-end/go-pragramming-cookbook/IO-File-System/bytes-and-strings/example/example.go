package main

import (
	"bytes_and_strings"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Example 1: bytes buffer")
	err := bytes_and_strings.BufferToStringExample("String -> Buffer -> Print!")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("\n\nExample 2: strings package example")
	s := "this is a test string for strings.*   "
	bytes_and_strings.SearchString(s)
	fmt.Println()
	bytes_and_strings.ModifyString(s)

	fmt.Println("\n\nExample 3: string -> io.Reader -> io.Writer")
	bytes_and_strings.StringReaderToWriter(s, os.Stdout)

}
