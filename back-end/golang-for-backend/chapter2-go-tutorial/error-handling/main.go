package main

import (
	"errors"
	"fmt"
)

var DivisionByZero = errors.New("division by zero")

func Divide(number, d float32) (float32, error) {
	if d == 0 {
		return 0, DivisionByZero
	}
	return number / d, nil
}

var SampleError = errors.New("This is a test error")

func recoverExample(invokeError bool) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Catch an error!")
			fmt.Println("error is.... ", err)
		} else {
			fmt.Println("No error")
		}
	}()

	if invokeError {
		panic(SampleError)
	}

	fmt.Println("Success!!")
}

func main() {
	fmt.Println("Custom Error Example")
	var num1, d1 float32 = 1, 1
	n1, e1 := Divide(num1, d1)
	if e1 != nil {
		fmt.Printf("%f / %f -> Error : %s\n", num1, d1, e1.Error())
	} else {
		fmt.Printf("%f / %f = %f\n", num1, d1, n1)
	}

	var num2, d2 float32 = 1, 0
	n2, e2 := Divide(num2, d2)
	if e2 != nil {
		fmt.Printf("%f / %f -> Error : %s\n", num2, d2, e2.Error())
	} else {
		fmt.Printf("%f / %f = %f\n", num2, d2, n2)
	}

	fmt.Println("\n\ntry-catch in golang")
	fmt.Println("[No Panic]")
	recoverExample(false)
	fmt.Println("[Panic]")
	recoverExample(true)
}
