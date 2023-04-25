package main

import (
	"fmt"
	"waitgroup"
)

func main() {
	sites := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://github.com",
		"https://google.com/search?q=golang",
	}

	result, err := waitgroup.Crawl(sites)
	if err != nil {
		panic(err)
	}
	fmt.Println("Crawling result:", result)
}
