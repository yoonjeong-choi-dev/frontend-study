package main

import (
	"log"
	"os"
)

func StandardOutLog() {
	l := log.New(os.Stdout, "standard output: ", log.Lshortfile)
	l.Println("logging to standard output")
}
