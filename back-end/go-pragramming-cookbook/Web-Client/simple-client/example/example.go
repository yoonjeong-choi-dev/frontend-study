package main

import (
	"client"
	"fmt"
)

const (
	google = "http://www.google.com"
	golang = "https://www.golang.org"
)

func main() {
	fmt.Println("Example 1. http.Client with secure: true, noOps: false")
	c := client.SetDefaultClient(true, false)
	controller := client.Controller{Client: c}
	if err := client.DoOperations(c, google); err != nil {
		panic(err)
	}

	if err := client.DefaultGetMethod(golang); err != nil {
		panic(err)
	}

	if err := controller.DoOperations(google); err != nil {
		panic(err)
	}
	if err := controller.DoOperations(golang); err != nil {
		panic(err)
	}

	fmt.Println("\nExample 2. http.Client with secure: true, noOps: true")
	client.SetDefaultClient(true, true)
	if err := client.DefaultGetMethod(golang); err != nil {
		panic(err)
	}
}
