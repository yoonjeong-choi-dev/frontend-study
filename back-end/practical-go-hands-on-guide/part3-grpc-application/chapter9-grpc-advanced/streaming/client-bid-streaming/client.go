package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	svc "streaming/service"
)

func setUpChat(r io.Reader, w io.Writer, client svc.UsersClient) error {
	stream, err := client.GetHelp(context.Background())
	if err != nil {
		return err
	}

	// User Setting
	user := svc.User{}
	for {
		scanner := bufio.NewScanner(r)
		prompt := "Name: "
		fmt.Fprint(w, prompt)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			return err
		}

		user.FirstName = scanner.Text()
		if user.FirstName == "" {
			fmt.Fprintln(w, "Please Enter your name")
		} else {
			break
		}
	}

	// chat streaming
	for {
		scanner := bufio.NewScanner(r)
		prompt := "Request: "
		fmt.Fprint(w, prompt)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			return err
		}

		msg := scanner.Text()
		if msg == "quit" {
			break
		}

		req := svc.UserHelpRequest{
			User:    &user,
			Request: msg,
		}
		err := stream.Send(&req)
		if err != nil {
			return err
		}

		res, err := stream.Recv()
		if err != nil {
			return err
		}

		fmt.Fprintf(w, "Response: %s\n", res.Response)
	}

	return stream.CloseSend()
}

func main() {
	clientConn, err := grpc.DialContext(
		context.Background(),
		":50051",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	defer clientConn.Close()

	if err != nil {
		log.Fatalf("Error for grpc connection: %v\n", err)
	}

	client := svc.NewUsersClient(clientConn)
	err = setUpChat(os.Stdin, os.Stdout, client)
	if err != nil {
		log.Fatal(err)
	}
}
