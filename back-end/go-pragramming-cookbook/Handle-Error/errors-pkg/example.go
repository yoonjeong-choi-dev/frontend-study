package main

import "fmt"

func main() {
	fmt.Println("Example 1. Wrapping Errors")
	WrapExample()

	fmt.Println("\nExample 2. Unwrapping Errors")
	UpWrapExample()

	fmt.Println("\nExample 3. StackTrace Errors")
	StackTraceExample()
}
