package main

import (
	"connectionpools"
	"fmt"
	"github.com/joho/godotenv"
	"runtime/debug"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Stack Trace:\n%s\n", debug.Stack())
		}
	}()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	if err := connectionpools.ExecWithTimeout(); err != nil {
		fmt.Printf("failed to exec: %s\n", err.Error())
		panic(err)
	}
}
