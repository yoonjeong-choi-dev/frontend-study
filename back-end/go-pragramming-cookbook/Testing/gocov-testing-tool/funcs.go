package gocov_testing_tool

import "fmt"

func exampleFunc1() error {
	fmt.Println("exampleFunc1")
	return nil
}

var exampleFunc2 = func() int {
	fmt.Println("exampleFunc2")
	return 7
}
