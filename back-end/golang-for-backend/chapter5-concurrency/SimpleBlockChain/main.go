package main

import (
	"fmt"
	"runtime"
)

func main() {
	prefix := "Tanmay Bakshi + Baheer Kamal"
	bitLength := 24

	fmt.Println("Single Thread Result")
	powSingleThread(prefix, bitLength)

	fmt.Printf("\nMulti Thread Result(with %d cores)\n", runtime.NumCPU())
	powMultiThread(prefix, bitLength)
}
