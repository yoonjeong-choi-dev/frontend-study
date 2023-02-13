package main

import "fmt"

func valueOfPi() float32 {
	return 3.141592
}

func multipleOfPi(multiplier uint) float32 {
	// unit -> float32
	return valueOfPi() * float32(multiplier)
}

func operateOnPi(multiplier uint, offset int) float32 {
	return (valueOfPi() - float32(offset)) * float32(multiplier)
}

func nameAndAge(uuid int) (string, int) {
	switch uuid {
	case 0:
		return "Yoonjeong", 31
	case 1:
		return "YJ", 29
	default:
		return "", -1
	}
}

func runMathOp(a int, b int, op func(int, int) int) int {
	return op(a, b)
}

func add(a int, b int) int { return a + b }
func sub(a int, b int) int { return a - b }
func mul(a int, b int) int { return a * b }
func div(a int, b int) int { return a / b }

func deferTest1(x int) int {
	defer fmt.Printf("Defer in deferTest1 - %d\n", x)

	y := x + 1
	fmt.Printf("deferTest1 - %d\n", y)
	return y
}

func deferTest2(x int) int {
	defer func() {
		fmt.Println("This is being called from an inline function")
		fmt.Println("Now we called another function containing defer")
		deferTest1(x)

		z := x - 1
		fmt.Printf("Defer in deferTest2 - %d\n", z)
	}()

	y := x + 1
	fmt.Printf("deferTest2 - %d\n", y)
	return y
}

func main() {
	fmt.Println("Basic Function")
	fmt.Println("Value of PI : ", valueOfPi())
	fmt.Printf("%d * PI = %f\n", 3, multipleOfPi(3))
	fmt.Printf("operatOnPi(%d, %d) = %f\n", 2, 1, operateOnPi(2, 1))

	fmt.Println("\n\nFunction with multiple return")
	var name1 string
	name1, age1 := nameAndAge(0)
	fmt.Printf("name: %s, age: %d\n", name1, age1)

	var name2 string
	var age2 int
	// cannot not use :=
	name2, age2 = nameAndAge(1)
	fmt.Printf("name: %s, age: %d\n", name2, age2)

	fmt.Println("\n\nFunction with function parameter")
	a := 12
	b := 5
	fmt.Printf("runMathOp(%d, %d, add) : %d\n", a, b, runMathOp(a, b, add))
	fmt.Printf("runMathOp(%d, %d, sub) : %d\n", a, b, runMathOp(a, b, sub))
	fmt.Printf("runMathOp(%d, %d, mul) : %d\n", a, b, runMathOp(a, b, mul))
	fmt.Printf("runMathOp(%d, %d, div) : %d\n", a, b, runMathOp(a, b, div))

	fmt.Println("\n\nFunction with defer 1")
	deferTest1(12)
	fmt.Println("\n\nFunction with defer 2")
	deferTest2(12)
}
