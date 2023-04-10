package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type TypedError struct {
	error
}

func WrappedError(e error) error {
	return errors.Wrap(e, "An error occurred in WrappedError")
}

func WrapExample() {
	e := errors.New("Standard error")
	fmt.Println("Standard Error:", WrappedError(e))

	typedErr := TypedError{errors.New("Typed error")}
	fmt.Println("Typed Error:", WrappedError(typedErr))

	fmt.Println("Nil Error:", WrappedError(nil))
}

func UpWrapExample() {
	err := error(TypedError{errors.New("sample error")})
	err = errors.Wrap(err, "wrapped")
	fmt.Println("wrapped err:", err)

	switch errors.Cause(err).(type) {
	case TypedError:
		fmt.Println("a typed error occurred:", err)
	default:
		fmt.Println("an unknown error occurred:", err)
	}
}

func StackTraceExample() {
	err := error(TypedError{errors.New("sample error")})
	err = errors.Wrap(err, "wrapped for stacktrace")
	fmt.Printf("Error Stack Trace: %+v\n", err)
	fmt.Println("EOF Error")
}
