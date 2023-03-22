package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"intercept/client/interceptor"
	svc "intercept/service"
	"io"
	"strconv"
	"time"
)

type userService struct {
	client svc.UsersClient
	reader io.Reader
	writer io.Writer
}

func (s *userService) InitClient(addr string) error {
	conn, err := grpc.DialContext(
		context.Background(),
		addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		// interceptor
		grpc.WithChainUnaryInterceptor(
			interceptor.MetaDataUnaryInterceptor,
			interceptor.LoggingUnaryInterceptor,
		),
		grpc.WithChainStreamInterceptor(
			interceptor.MetaDataStreamInterceptor,
			interceptor.LoggingStreamInterceptor,
		),
	)

	if err != nil {
		return err
	}

	s.client = svc.NewUsersClient(conn)
	return nil
}

func (s *userService) InitInteraction(r io.Reader, w io.Writer) {
	s.reader = r
	s.writer = w
}

// GetUser Unary Pattern
func (s *userService) GetUser(req *svc.UserGetRequest) (*svc.User, error) {
	return s.client.GetUser(context.Background(), req)
}

func (s *userService) GetUserInteraction() {
	fmt.Fprintln(s.writer, "GetUser Service")

	input, err := getUserInputString(s.reader, s.writer, "Full Name")
	if checkError(s.writer, "Scanner", err) {
		return
	}

	req := svc.UserGetRequest{
		Id:       strconv.FormatInt(time.Now().Unix(), 10),
		FullName: input,
	}

	res, err := s.GetUser(&req)
	printResponse(s.writer, res, err)
}

// CreateUser Server-Side Stream
func (s *userService) CreateUser(req *svc.User) ([]*svc.UserCreateLog, error) {
	stream, err := s.client.CreateUser(context.Background(), req)
	if err != nil {
		return nil, err
	}

	var ret []*svc.UserCreateLog
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return ret, err
		}

		ret = append(ret, res)
	}

	return ret, nil
}

func (s *userService) CreateUserInteraction() {
	fmt.Fprintln(s.writer, "CreateUser Service")

	// User Setting
	user, err := s.getUserByInput()
	if checkError(s.writer, "Scanner", err) {
		return
	}

	res, err := s.CreateUser(user)
	printResponse(s.writer, res, err)
}

// GetUsers Client-Side Stream
func (s *userService) GetUsers(users []*svc.User) (*svc.UsersList, error) {
	stream, err := s.client.GetUsers(context.Background())
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		err := stream.Send(user)
		if err != nil {
			return nil, err
		}
	}
	return stream.CloseAndRecv()
}

func (s *userService) GetUsersInteraction() {
	fmt.Fprintln(s.writer, "GetUsers Service")

	numUsers, err := getUserInputInt(s.reader, s.writer, "Number of Users")
	if checkError(s.writer, "Scanner", err) {
		return
	}

	var users []*svc.User
	for i := 0; i < numUsers; i++ {
		fmt.Fprintln(s.writer, "Enter User Info(Enter 'quit' to finish)")
		user, err := s.getUserByInput()
		if checkError(s.writer, "Scanner", err) {
			return
		}
		users = append(users, user)
	}

	res, err := s.GetUsers(users)
	printResponse(s.writer, res, err)
}

// GetHelpInteraction Bidirectional Stream
func (s *userService) GetHelpInteraction() error {
	stream, err := s.client.GetHelp(context.Background())
	if err != nil {
		return err
	}

	// User Setting
	user, err := s.getUserByInput()
	if checkError(s.writer, "Scanner", err) {
		return err
	}

	// chat streaming
	for {
		input, err := getUserInputString(s.reader, s.writer, "Request")
		if input == "quit" {
			break
		}

		req := svc.UserHelpRequest{
			User:    user,
			Request: input,
		}

		if err = stream.Send(&req); err != nil {
			return err
		}

		res, err := stream.Recv()
		if err != nil {
			return err
		}

		fmt.Fprintf(s.writer, "Response: %s\n", res.Response)
	}

	return stream.CloseSend()
}

func (s *userService) getUserByInput() (*svc.User, error) {
	// User Setting
	user := svc.User{
		Id: strconv.FormatInt(time.Now().Unix(), 10),
	}

	input, err := getUserInputString(s.reader, s.writer, "First Name")
	if err != nil {
		return nil, err
	}
	user.FirstName = input

	input, err = getUserInputString(s.reader, s.writer, "Last Name")
	if err != nil {
		return nil, err
	}
	user.LastName = input

	return &user, nil
}
