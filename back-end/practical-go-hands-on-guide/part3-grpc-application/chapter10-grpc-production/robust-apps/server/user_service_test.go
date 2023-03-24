package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"io"
	"net"
	svc "robust-app/service"
	"testing"
)

func createUserServiceClient(l *bufconn.Listener) (svc.UsersClient, error) {
	dialer := func(ctx context.Context, addr string) (net.Conn, error) {
		return l.Dial()
	}

	client, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithInsecure(),
		grpc.WithContextDialer(dialer),
	)
	if err != nil {
		return nil, err
	}
	return svc.NewUsersClient(client), nil
}

func TestUserService_GetUser(t *testing.T) {
	l, _ := startTestServer()

	client, err := createUserServiceClient(l)
	checkError("Creating Dial", err)

	id := "test-id"
	fullName := "yoonjeong choi"
	res, err := client.GetUser(
		context.Background(),
		&svc.UserGetRequest{
			Id:       id,
			FullName: fullName,
		},
	)
	checkError("Response", err)
	checkStringTypeResponse(t, "Id", id, res.Id)
	checkStringTypeResponse(t, "FullName", fullName,
		fmt.Sprintf("%s %s", res.FirstName, res.LastName),
	)
	printResponse(t, res)
}

func TestUserService_CreateUser(t *testing.T) {
	l, _ := startTestServer()

	client, err := createUserServiceClient(l)
	checkError("Creating Dial", err)

	req := svc.User{
		Id:        "test-id",
		FirstName: "Yoonjeong",
		LastName:  "Choi",
	}

	stream, err := client.CreateUser(
		context.Background(),
		&req,
	)
	checkError("Response", err)

	var reps []*svc.UserCreateLog
	for {
		userLog, err := stream.Recv()

		if err == io.EOF {
			break
		}
		checkError("Reading Stream", err)

		reps = append(reps, userLog)
	}

	checkIntTypeResponse(t, "Number of Response", 5, len(reps))

	var expectedLog string
	for idx, userLog := range reps {
		if idx == 0 {
			expectedLog = fmt.Sprintf("Starting Creating User(%s) For %s-%s",
				req.Id,
				req.FirstName,
				req.LastName,
			)
		} else if idx == 4 {
			expectedLog = fmt.Sprintf("Success to Create User(%s) For %s-%s",
				req.Id,
				req.FirstName,
				req.LastName,
			)
		} else {
			expectedLog = fmt.Sprintf("[Step %d] Save to Repostory %d-%s",
				idx,
				idx,
				req.FirstName,
			)
		}

		checkStringTypeResponse(t, "Log", expectedLog, userLog.Log)
		if err := userLog.Timestamp.CheckValid(); err != nil {
			t.Errorf("Invalid Timestamp: %v\n", err)
		}
	}
	printResponse(t, reps)
}

func TestUserService_GetUsers(t *testing.T) {
	l, _ := startTestServer()

	client, err := createUserServiceClient(l)
	checkError("Creating Dial", err)

	// test data
	var reqUsers []*svc.User
	reqSize := 3
	for i := 0; i < reqSize; i++ {
		reqUsers = append(reqUsers, &svc.User{
			Id:        "test-id",
			FirstName: "Yoonjeong",
			LastName:  "Choi",
		})
	}

	stream, err := client.GetUsers(context.Background())
	for _, req := range reqUsers {
		err = stream.Send(req)
		checkError("Sending Stream", err)
	}

	res, err := stream.CloseAndRecv()
	resUsers := res.Users
	checkIntTypeResponse(t, "Number of Response", reqSize, len(resUsers))

	for idx, user := range resUsers {
		checkStringTypeResponse(t, "Id",
			fmt.Sprintf("[Req #%d]%s", idx, reqUsers[idx].Id),
			user.Id,
		)
	}

	printResponse(t, res)
}
