package main

import "fmt"

func calculateSquare(inputChan chan int, outputChan chan []int) {
	for x := range inputChan {
		outputChan <- []int{x, x * x}
	}
}

func main() {
	inputChan := make(chan int)

	const maxLoop = 10
	outputChan := make(chan []int, maxLoop)

	go calculateSquare(inputChan, outputChan)

	for i := 0; i < maxLoop; i++ {
		inputChan <- i
	}
	close(inputChan)

	for i := 0; i < maxLoop; i++ {
		ret := <-outputChan
		fmt.Printf("pow(%d, 2) = %d\n", ret[0], ret[1])
	}

}

func deadlockCode() {
	inputChan := make(chan int)
	outputChan := make(chan []int)

	go calculateSquare(inputChan, outputChan)

	for i := 0; i < 10; i++ {
		inputChan <- i
	}

	for ret := range outputChan {
		fmt.Printf("pow(%d, 2) = %d\n", ret[0], ret[1])
	}
}
