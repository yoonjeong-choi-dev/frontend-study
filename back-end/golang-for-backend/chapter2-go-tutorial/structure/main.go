package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func printUser(user User) {
	fmt.Printf("Name: %s, Age: %d\n", user.Name, user.Age)
}

func wrongIncrementAge(user User) {
	user.Age++
}

func incrementAge(user *User) {
	user.Age++
}

func (user *User) prettyString() string {
	return fmt.Sprintf("%s is %d years old!", user.Name, user.Age)
}

func (user *User) incrementAge() {
	user.Age++
}

func main() {
	fmt.Println("Simple Struct")
	user1 := User{"Yoonjeong", 30}
	printUser(user1)

	fmt.Print("wrongIncrementAge -> ")
	wrongIncrementAge(user1)
	printUser(user1)

	fmt.Print("incrementAge -> ")
	incrementAge(&user1)
	printUser(user1)

	fmt.Println("\n\nStruct with its methods")
	user2 := User{"YJ", 29}
	fmt.Println("To String : ", user2.prettyString())

	user2.incrementAge()
	fmt.Print("user2.incrementAge -> ", user2.prettyString())
}
