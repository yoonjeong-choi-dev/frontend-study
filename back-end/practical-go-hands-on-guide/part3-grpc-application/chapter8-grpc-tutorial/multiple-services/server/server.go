package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	svc "multiple-services/service"
	"net"
	"strings"
)

type userService struct {
	svc.UnimplementedUsersServer
}

type repoService struct {
	svc.UnimplementedRepoServer
}

func (s *userService) GetUser(
	ctx context.Context, in *svc.UserGetRequest) (*svc.UserGetResponse, error) {
	log.Printf("Recieved request for user with Email: %s, Id: %s\n",
		in.Email, in.Id,
	)

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

func (s *repoService) GetRepos(
	ctx context.Context, in *svc.RepoGetRequest) (*svc.RepoGetResponse, error) {
	log.Printf("Recieved request for repo with Id: %s, CreatedId: %s\n",
		in.Id, in.CreatorId,
	)

	repo := svc.Repository{
		Id:   in.Id,
		Name: "moloco-study",
		Url:  "github.com/yoonjeong-choi-dev",
		Owner: &svc.User{
			Id:        in.CreatorId,
			FirstName: "Yoonjeong",
			LastName:  "Choi",
			Age:       31,
		},
	}

	res := svc.RepoGetResponse{
		Repo: []*svc.Repository{&repo},
	}
	return &res, nil
}

func registerService(s *grpc.Server) {
	svc.RegisterUsersServer(s, &userService{})
	svc.RegisterRepoServer(s, &repoService{})
}

func startServer(s *grpc.Server, l net.Listener) error {
	return s.Serve(l)
}

func main() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error for TCP Listener: %v\n", err)
	}

	s := grpc.NewServer()
	registerService(s)
	log.Fatal(startServer(s, l))
}
