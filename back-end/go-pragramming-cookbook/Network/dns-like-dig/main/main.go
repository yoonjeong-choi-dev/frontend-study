package main

import (
	dns_like_dig "dns-like-dig"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <address>\n", os.Args[0])
		os.Exit(1)
	}

	address := os.Args[1]
	lookup, err := dns_like_dig.LookupAddress(address)
	if err != nil {
		fmt.Errorf("failed to lookup: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(lookup)
}
