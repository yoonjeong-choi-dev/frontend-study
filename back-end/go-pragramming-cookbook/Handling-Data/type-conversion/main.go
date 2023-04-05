package main

import "fmt"

func main() {
	fmt.Println("Example 1. Basic Type Conversion")
	BasicTypeConversion()

	fmt.Println("\nExample 2. String Conversion with strconv")
	if err := StringConvert(); err != nil {
		fmt.Println("String Conversion Error", err)
	}

	fmt.Println("\nExample 3. Conversion with reflect")
	InterfaceConversion()
}
