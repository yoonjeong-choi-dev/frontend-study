package main

import (
	"fmt"
	"runtime/debug"
	"strconv"
)

func RisePanic() {
	zero, err := strconv.ParseInt("0", 10, 64)
	if err != nil {
		panic(err)
	}

	div := 1 / zero
	fmt.Printf("Cannot reach this code - %d\n", div)
}

func CatchPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[Panic occurred]\n", r)
			fmt.Println("[Stacktrace from panic]\n" + string(debug.Stack()))
		}
	}()

	RisePanic()
}
