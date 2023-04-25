package main

import (
	context_example "context-example"
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Try %d: ", i)
		context_example.Example()
	}
}
