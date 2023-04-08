package main

import (
	"fmt"
	null_encoding "null-enocding"
)

func main() {
	fmt.Println("Example 1. Basic Json encoding and decoding")
	if err := null_encoding.BaseEncoding(); err != nil {
		panic(err)
	}

	fmt.Println("\nExample 2. Pointer Json encoding and decoding")
	if err := null_encoding.PointerExample(); err != nil {
		panic(err)
	}

	fmt.Println("\nExample 3. Custom Null Type Json encoding and decoding")
	if err := null_encoding.CustomNullTypeExample(); err != nil {
		panic(err)
	}
}
