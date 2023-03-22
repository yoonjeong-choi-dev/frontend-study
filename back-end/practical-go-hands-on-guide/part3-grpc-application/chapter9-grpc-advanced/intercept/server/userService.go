package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	svc "intercept/service"
	"intercept/utils"
	"io"
	"log"
	"strings"
	"time"
)

type userService struct {
	svc.UnimplementedUsersServer
}

// GetUser Unary Pattern
func (s *userService) GetUser(
	ctx context.Context, in *svc.UserGetRequest,
) (*svc.User, error) {
	names := strings.Split(in.FullName, " ")
	if len(names) < 2 {
		return nil, status.Error(
			codes.InvalidArgument,
			"please enter your full name with space separator",
		)
	}

	user := svc.User{
		Id:        in.Id,
		FirstName: names[0],
		LastName:  names[len(names)-1],
	}

	return &user, nil
}

// CreateUser Server-Side Stream
func (s *userService) CreateUser(
	in *svc.User,
	stream svc.Users_CreateUserServer,
) error {

	userLog := svc.UserCreateLog{
		Log: fmt.Sprintf("Starting Creating User(%s) For %s-%s",
			in.Id,
			in.FirstName,
			in.LastName,
		),
		Timestamp: timestamppb.Now(),
	}

	if err := stream.Send(&userLog); err != nil {
		return err
	}

	step := 1
	for {
		userLog.Log = fmt.Sprintf("[Step %d] Save to Repostory %d-%s", step, step, in.FirstName)
		userLog.Timestamp = timestamppb.Now()

		if err := stream.Send(&userLog); err != nil {
			return err
		}
		if step >= 3 {
			break
		}

		step++
		time.Sleep(300 * time.Millisecond)
	}

	userLog.Log = fmt.Sprintf("Success to Create User(%s) For %s-%s",
		in.Id,
		in.FirstName,
		in.LastName,
	)
	userLog.Timestamp = timestamppb.Now()
	if err := stream.Send(&userLog); err != nil {
		return err
	}
	return nil
}

// GetUsers Client-Side Stream
func (s *userService) GetUsers(stream svc.Users_GetUsersServer) error {
	var res []*svc.User

	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		req.Id = fmt.Sprintf("[Req #%d]%s", count, req.Id)
		res = append(res, req)
		count++
	}
	return stream.SendAndClose(&svc.UsersList{
		Users: res,
	})
}

// GetHelp Bidirectional Stream
func (s *userService) GetHelp(stream svc.Users_GetHelpServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error for reading streaming: %s\n", utils.GetJsonStringUnsafe(err))
			return err
		}

		log.Printf("Request Recieved Message: %s\n", req.Request)
		res := svc.UserHelpResponse{
			Response: fmt.Sprintf("Hi~ %s!. Your help message is '%s'",
				req.User.FirstName,
				req.Request,
			),
		}

		err = stream.Send(&res)
		if err != nil {
			log.Printf("Error for sending streaming: %s\n", utils.GetJsonStringUnsafe(err))
			return err
		}
	}

	return nil
}
