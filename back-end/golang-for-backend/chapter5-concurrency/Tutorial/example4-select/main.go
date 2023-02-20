package main

import "fmt"

func calculateSquare(inputChan chan int, outputChan chan []int, exitChan chan int) {
	var input int
	for {
		select {
		case input = <-inputChan:
			outputChan <- []int{input, input * input}
		case <-exitChan:
			return
		}
	}
}

func main() {
	inputChan := make(chan int)

	const maxLoop = 10
	outputChan := make(chan []int, maxLoop)

	exitChan := make(chan int)

	go calculateSquare(inputChan, outputChan, exitChan)

	for i := 0; i < maxLoop; i++ {
		inputChan <- i
	}

	for i := 0; i < maxLoop; i++ {
		ret := <-outputChan
		fmt.Printf("pow(%d, 2) = %d\n", ret[0], ret[1])
	}

	fmt.Println("Try to terminate the go-routine function")
	exitChan <- 1
	fmt.Println("Exit the program")
}
