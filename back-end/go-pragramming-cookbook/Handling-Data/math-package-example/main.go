package main

import "fmt"

func main() {
	fmt.Println("Example for math package")
	mathPackageExample()

	fmt.Println("\nFibonacci Values")
	for i := 0; i < 100; i++ {
		fmt.Printf("%v ", Fib(i))
	}
	fmt.Println()
}
