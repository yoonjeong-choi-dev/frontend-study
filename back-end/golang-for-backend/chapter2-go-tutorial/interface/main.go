package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (person *Person) incrementAge() {
	person.Age++
}

func (person *Person) getAge() int {
	return person.Age
}

type Dog struct {
	Name  string
	Age   int
	Owner *Person
}

func (dog *Dog) incrementAge() {
	dog.Age++
}

func (dog *Dog) getAge() int {
	return dog.Age
}

type Living interface {
	incrementAge()
	getAge() int
}

func incrementAgeAndPrintAge(being Living) {
	fmt.Printf("Before: %d", being.getAge())
	being.incrementAge()
	fmt.Printf(" -> After: %d\n", being.getAge())
}

func main() {
	p1 := Person{"YJ", 31}
	d1 := Dog{"Dog", 5, &p1}

	incrementAgeAndPrintAge(&p1)
	incrementAgeAndPrintAge(&d1)
}
