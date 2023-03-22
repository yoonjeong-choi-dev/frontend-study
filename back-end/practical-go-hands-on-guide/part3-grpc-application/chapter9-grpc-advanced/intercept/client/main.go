package main

import (
	"fmt"
	svc "intercept/service"
	"intercept/utils"
	"os"
)

func printUsage() {
	fmt.Printf("\n\n")
	fmt.Println("========================================================")
	fmt.Println("1.Unary Pattern - GetUser Interaction")
	fmt.Println("2.Server Stream Pattern - CreateUser Interaction")
	fmt.Println("3.Client Stream Pattern - GetUsers Interaction")
	fmt.Println("4.Bidirectional Pattern - GetHelp Interaction")
	fmt.Println("5.Simple Examples")
	fmt.Println("Otherwise: Quit the program")
	fmt.Println("========================================================")
}

func examples(client userService) {
	fmt.Println("========================================================")
	fmt.Println("GetUser Examples")
	getReq := svc.UserGetRequest{
		Id:       "create-test",
		FullName: "Yoonjeong Dev",
	}
	getRes, err := client.GetUser(&getReq)
	printResponse(os.Stdout, getRes, err)

	fmt.Println("\n\n========================================================")
	fmt.Println("CreateUser Examples")
	user := svc.User{
		Id:        "Test User",
		FirstName: "Yoonjeong",
		LastName:  "Choi",
	}
	createRes, err := client.CreateUser(&user)
	printResponse(os.Stdout, createRes, err)

	fmt.Println("\n\n========================================================")
	fmt.Println("GetUsers Examples")
	var reqUsers []*svc.User
	reqSize := 3
	for i := 0; i < reqSize; i++ {
		reqUsers = append(reqUsers, &svc.User{
			Id:        "test-id",
			FirstName: "Yoonjeong",
			LastName:  "Choi",
		})
	}

	usersRes, err := client.GetUsers(reqUsers)
	printResponse(os.Stdout, usersRes, err)
}

func main() {
	client := userService{}
	client.InitClient("localhost:50051")
	client.InitInteraction(os.Stdin, os.Stdout)

	for {
		printUsage()

		option, err := getUserInputInt(os.Stdin, os.Stdout, "Select the option: ")
		if checkError(os.Stdout, "Scanner", err) {
			return
		}

		fmt.Println()
		switch option {
		case 1:
			client.GetUserInteraction()
		case 2:
			client.CreateUserInteraction()
		case 3:
			client.GetUsersInteraction()
		case 4:
			err := client.GetHelpInteraction()
			if err != nil {
				fmt.Printf("GetHelpInteraction Error :%s\n", utils.GetJsonStringUnsafe(err.Error()))
			}
		case 5:
			examples(client)
		default:
			fmt.Println("Quit the Program")
			return
		}
	}
}
