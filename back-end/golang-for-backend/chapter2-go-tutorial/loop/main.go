package main

import "fmt"

func main() {
	names := []string{"Choi", "Yoonjeong", "Yoonjeong Choi"}

	fmt.Println("for-in loop : Index")
	for i := range names {
		fmt.Printf("[%d] names[i] = %s\n", i, names[i])
	}

	fmt.Println("for-in loop :Only Values")
	for _, v := range names {
		fmt.Println(v)
	}

	fmt.Println("for-in loop :(index, value)")
	for i, v := range names {
		fmt.Printf("[%d] %s\n", i, v)
	}

	fmt.Println("C-style classic loop")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	fmt.Println("Classic Loop like While")
	i := 1
	for i < 1000 {
		i += i
	}
	fmt.Printf("Power of 2 >= 1000 : %d\n", i)
}
