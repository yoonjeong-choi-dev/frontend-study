package main

import "fmt"

type CustomError struct {
	Result string
	Type   int
}

func (c CustomError) Error() string {
	return fmt.Sprintf("there is an error(type %d); %s was the result", c.Type, c.Result)
}

func CustomErrorExample() {
	err := CustomError{
		Result: "'cannot handle'",
		Type:   3,
	}

	fmt.Println("Custom error: ", err)
}
