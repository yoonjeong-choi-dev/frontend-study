package main

import (
	"fmt"
	various_type_encoding "various-type-encoding"
)

func main() {
	fmt.Println("Example 1. Gob Encoding and Decoding")
	if err := various_type_encoding.GobExample(); err != nil {
		panic(err)
	}

	fmt.Println("\nExample 2. Base64 URL Encoding and Decoding")
	if err := various_type_encoding.Base64Example(); err != nil {
		panic(err)
	}

	fmt.Println("\nExample 3. Base64 Standard Encoding and Decoding using encoder&decoder")
	if err := various_type_encoding.Base64ExampleWithEncoder(); err != nil {
		panic(err)
	}
}
