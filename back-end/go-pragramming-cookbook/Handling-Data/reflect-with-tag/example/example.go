package main

import (
	"fmt"
	reflect_with_tag "reflect-with-tag"
)

func main() {
	if err := reflect_with_tag.EmptyStructExample(); err != nil {
		panic(err)
	}

	fmt.Println()

	if err := reflect_with_tag.FullStructExample(); err != nil {
		panic(err)
	}
}
