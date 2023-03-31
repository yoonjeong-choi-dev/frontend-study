package main

import (
	"cli-with-flag/config"
	"flag"
	"fmt"
)

func main() {
	c := config.Config{}
	c.Setup()

	flag.Parse()
	fmt.Printf("Your Message is..\n%s\n", c.GetMessage())
}
