package main

import "fmt"

func main() {
	age := 31
	if age >= 18 {
		fmt.Println("Welcome!")
	} else {
		fmt.Println("You are too young")
	}

	name := "Yoonjeong Choi"
	switch name {
	case "YJ":
		fmt.Println("Hello YJ!")
	case "Yoonjeong Choi":
		fmt.Println("Hello Yoonjeong~")
	case "Test":
		fmt.Println("This is test")
	default:
		fmt.Println("Who are you..?")
	}
}
