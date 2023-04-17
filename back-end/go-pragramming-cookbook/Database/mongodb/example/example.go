package main

import (
	"fmt"
	"mongodb"
	"runtime/debug"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Printf("Stack Trace:\n%s\n", debug.Stack())
		}
	}()

	if err := mongodb.MongoExample("mongodb://localhost"); err != nil {
		panic(err)
	}
}
