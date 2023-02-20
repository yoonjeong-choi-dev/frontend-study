package main

import (
	"fmt"
	"time"
)

func main() {
	// 2번째 인자로 숫자를 넣어주는 경우 세마포어와 비슷한 역할 가능
	mutex := make(chan struct{})
	go func() {
		fmt.Println("Signalling")
		time.Sleep(1000 * time.Millisecond)
		mutex <- struct{}{}

		fmt.Println("End of Sub Function")
	}()
	fmt.Println("After calling the Sub Function")

	<-mutex
	fmt.Println("Exit the program")
}
