package main

import "fmt"

func main() {
	fmt.Println("Before Panic")
	CatchPanic()
	fmt.Println("After Panic")
}
