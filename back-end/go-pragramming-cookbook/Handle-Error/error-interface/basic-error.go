package main

import (
	"errors"
	"fmt"
)

// 패키지 수준의 에러
var ErrorValue = errors.New("this is a custom error variable")

type TypedError struct {
	error
}

func BasicErrorExample() {
	err := errors.New("this is a new error")
	fmt.Println("errors.New: ", err)

	err = fmt.Errorf("error with formating(%s)", "some error")
	fmt.Println("fmt.Errorf: ", err)

	err = ErrorValue
	fmt.Println("value error: ", err)

	err = TypedError{errors.New("error with wrapping")}
	fmt.Println("Error wrapping struct: ", err)
}
