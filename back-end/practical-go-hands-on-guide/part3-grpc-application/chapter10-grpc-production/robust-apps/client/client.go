package main

import (
	"fmt"
	"os"
	svc "robust-app/service"
	"robust-app/utils"
)

func printUsage() {
	fmt.Printf("\n\n")
	fmt.Println("========================================================")
	fmt.Println("1.Unary Pattern - GetUser Interaction")
	fmt.Println("2.Server Stream Pattern - CreateUser Interaction")
	fmt.Println("3.Client Stream Pattern - GetUsers Interaction")
	fmt.Println("4.Bidirectional Pattern - GetHelp Interaction")
	fmt.Println("5.Simple Examples")
	fmt.Println("===============================")
	fmt.Println("6.Unary Panic Test")
	fmt.Println("7.Server Stream Panic Test")
	fmt.Println("Otherwise: Quit the program")
	fmt.Println("========================================================")
}

func examples(userClient userService, healthClient healthService) {
	fmt.Println("========================================================")
	fmt.Println("GetUser Examples")
	if healthClient.GetUserServiceStatus() {
		fmt.Println("Cannot Connect With Server")
		os.Exit(1)
	}

	getReq := svc.UserGetRequest{
		Id:       "create-test",
		FullName: "Yoonjeong Dev",
	}
	getRes, err := userClient.GetUser(&getReq)
	printResponse(os.Stdout, getRes, err)

	fmt.Println("\n\n========================================================")
	fmt.Println("CreateUser Examples")
	if healthClient.GetUserServiceStatus() {
		fmt.Println("Cannot Connect With Server")
		os.Exit(1)
	}

	user := svc.User{
		Id:        "Test User",
		FirstName: "Yoonjeong",
		LastName:  "Choi",
	}
	createRes, err := userClient.CreateUser(&user)
	printResponse(os.Stdout, createRes, err)

	fmt.Println("\n\n========================================================")
	fmt.Println("GetUsers Examples")
	if healthClient.GetUserServiceStatus() {
		fmt.Println("Cannot Connect With Server")
		os.Exit(1)
	}

	var reqUsers []*svc.User
	reqSize := 3
	for i := 0; i < reqSize; i++ {
		reqUsers = append(reqUsers, &svc.User{
			Id:        "test-id",
			FirstName: "Yoonjeong",
			LastName:  "Choi",
		})
	}

	usersRes, err := userClient.GetUsers(reqUsers)
	printResponse(os.Stdout, usersRes, err)
}

func main() {
	addr := "localhost:50051"
	userClient := userService{}
	userClient.InitInteraction(os.Stdin, os.Stdout)
	connErr := userClient.InitClient(addr)
	if checkError(os.Stdout, "Connection for UserService", connErr) {
		os.Exit(1)
	}
	defer userClient.conn.Close()
	defer userClient.cancel()

	healthClient := healthService{}
	healthClient.InitInteraction(os.Stdin, os.Stdout)
	connErr = healthClient.InitClient(addr)
	if checkError(os.Stdout, "Connection for HealthCheck", connErr) {
		os.Exit(1)
	}
	defer healthClient.conn.Close()
	defer healthClient.cancel()

	for {
		printUsage()

		option, err := getUserInputInt(os.Stdin, os.Stdout, "Select the option: ")
		if checkError(os.Stdout, "Scanner", err) {
			return
		}

		if healthClient.GetUserServiceStatus() {
			fmt.Println("Cannot Connect With Server")
			os.Exit(1)
		}

		fmt.Println()
		switch option {
		case 1:
			userClient.GetUserInteraction()
		case 2:
			userClient.CreateUserInteraction()
		case 3:
			userClient.GetUsersInteraction()
		case 4:
			err := userClient.GetHelpInteraction()
			if err != nil {
				fmt.Printf("GetHelpInteraction Error :%s\n", utils.GetJsonStringUnsafe(err.Error()))
			}
		case 5:
			examples(userClient, healthClient)
		case 6:
			userClient.RisePanicTest()
		case 7:
			userClient.RiseServerStreamPanic()
		default:
			fmt.Println("Quit the Program")
			return
		}
	}
}
