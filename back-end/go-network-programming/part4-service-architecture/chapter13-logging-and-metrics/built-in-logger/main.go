package main

import "fmt"

func main() {
	fmt.Println("Example 1. Standard Output Logger")
	StandardOutLog()

	fmt.Println("\nExample 2. Multi Writer Logger")
	MultiWriterLog()

	fmt.Println("\nExample 3. Implement Simple Log Level Logger")
	LogLevelExample()
}
