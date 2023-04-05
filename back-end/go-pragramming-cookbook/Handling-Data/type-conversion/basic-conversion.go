package main

import "fmt"

func BasicTypeConversion() {
	var integer = 24
	var float = 2.0

	// int -> float
	multi := float64(integer) * float

	// float -> string
	fToStr := fmt.Sprintf("%.2f", multi)

	fmt.Printf("integer: %d (type:%T)\n", integer, integer)
	fmt.Printf("float: %f (type:%T)\n", multi, multi)
	fmt.Printf("float to string: %s (type:%T)\n", fToStr, fToStr)
}
