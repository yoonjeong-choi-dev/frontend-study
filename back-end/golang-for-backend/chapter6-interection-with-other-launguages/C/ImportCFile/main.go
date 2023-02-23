package main

//#cgo LDFLAGS: -L. -lyjfunction
//void PrintHelloWorld();
//int Factorial(int x);
import "C"
import "fmt"

func main() {
	C.PrintHelloWorld()

	fmt.Println("\nFactorial Test")
	for i := 1; i < 8; i++ {
		ret := C.Factorial(C.int(i))
		fmt.Printf("%d! = %d\n", i, ret)
	}
}
