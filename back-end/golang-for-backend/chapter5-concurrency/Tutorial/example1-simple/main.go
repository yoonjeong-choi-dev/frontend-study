package main

import (
	"fmt"
	"time"
)

func printSquare(x int) {
	fmt.Println("==========================")
	fmt.Println("This is an expansive call")

	time.Sleep(1 * time.Millisecond)
	fmt.Printf("pow(%d, 2) = %d\n", x, x*x)

	fmt.Println("==========================")
}

func main() {
	go printSquare(7)

	fmt.Println("Next line of expansive function call!")

	// printSquare 관련 고루틴이 끝나기 전에 main 고루틴이 종료되기 때문에 잠시 main 고루틴 대기
	// : 임시 방편으로, 실제 고루틴이 끝날 때까지 기다리려면 채널이 필요
	time.Sleep(5 * time.Millisecond)
}
