package main

import "fmt"

var age = 12
var ageNotInit int

var (
	name  = "Yoonjeong"
	myAge = 31
)

const ageConst = 31

func main() {
	fmt.Println("\n\nGlobal Variables")
	fmt.Printf("age: %d\n", age)
	ageNotInit = 24
	fmt.Printf("ageNotInit: %d\n", ageNotInit)

	age = 31
	fmt.Printf("After re-assigning age: %d\n", age)

	fmt.Printf("Name : %s, Age: %d\n", name, myAge)

	fmt.Printf("Constant : %d\n", ageConst)
	//ageConst = 31

	fmt.Println("\n\nLocal Variables")
	nameLoc := "Local Name"
	fmt.Printf("nameLoc : %s\n", nameLoc)

	fmt.Println("\n\nArray")
	simpleArr := []string{"Name1", "Name2"}
	fmt.Println("Simple Array :", simpleArr)
	simpleArr = append(simpleArr, "Append Name")
	fmt.Println("Simple Array after append:", simpleArr)

	makeArr := make([]string, 3)
	makeArr[0] = "Name 1"
	makeArr[1] = "Name 2"
	makeArr[2] = "Last Name"
	fmt.Println("Array with make :", makeArr)

	capacity := 3
	arrWithCapacity := make([]int, 0, capacity)
	fmt.Printf("current len : %d\n", len(arrWithCapacity))
	fmt.Println(arrWithCapacity)

	arrWithCapacity = append(arrWithCapacity, 12)
	fmt.Printf("current len : %d\n", len(arrWithCapacity))
	fmt.Println(arrWithCapacity)

	arrWithCapacity = append(arrWithCapacity, 24)
	fmt.Printf("current len : %d\n", len(arrWithCapacity))
	fmt.Println(arrWithCapacity)

	arrWithCapacity = append(arrWithCapacity, 36)
	fmt.Printf("current len : %d\n", len(arrWithCapacity))
	fmt.Println(arrWithCapacity)

}
