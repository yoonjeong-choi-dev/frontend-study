package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	svc "streaming/service"
	"streaming/utils"
	"strings"
)

type userService struct {
	svc.UnimplementedUsersServer
}

// GetUser Unary Pattern
func (s *userService) GetUser(
	ctx context.Context, in *svc.UserGetRequest) (*svc.UserGetResponse, error) {
	components := strings.Split(in.Email, "@")
	if len(components) != 2 {
		return nil, status.Error(
			codes.InvalidArgument, "invalid email address",
		)
	}

	user := svc.User{
		Id:        in.Id,
		FirstName: components[0],
		LastName:  components[1],
		Age:       31,
	}
	return &svc.UserGetResponse{User: &user}, nil
}

// GetHelp Bi-Directional Stream
func (s *userService) GetHelp(stream svc.Users_GetHelpServer) error {
	log.Println("Start Listening Client Streaming Data")

	// 스트리밍 연결이 끊길 떄까지 반복
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
			Response: fmt.Sprintf("Hi~ %s! Help for %s", req.User.FirstName, req.Request),
		}

		err = stream.Send(&res)
		if err != nil {
			log.Printf("Error for sending streaming: %s\n", utils.GetJsonStringUnsafe(err))
			return err
		}

	}

	log.Println("Finished Listening Client Streaming Data")
	return nil
}
