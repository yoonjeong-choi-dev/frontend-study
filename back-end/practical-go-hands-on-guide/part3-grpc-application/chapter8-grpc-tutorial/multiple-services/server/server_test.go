package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	svc "multiple-services/service"
	"net"
	"testing"
)

func startTestGrpcServer() (*grpc.Server, *bufconn.Listener) {
	l := bufconn.Listen(10)
	s := grpc.NewServer()
	registerService(s)

	go func() {
		err := startServer(s, l)
		if err != nil {
			log.Fatalf("Error starting server: %v\n", err)
		}
	}()
	return s, l
}

func TestRepoService_GetRepos(t *testing.T) {
	s, l := startTestGrpcServer()
	defer s.GracefulStop()

	bufconnDialer := func(ctx context.Context, addr string) (net.Conn, error) {
		return l.Dial()
	}

	client, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithInsecure(),
		grpc.WithContextDialer(bufconnDialer),
	)
	if err != nil {
		t.Fatalf("Error for creating grpc client: %v\n", err)
	}

	repoClient := svc.NewRepoClient(client)

	id := "repo-id"
	creatorId := "yoonjeong-id"
	res, err := repoClient.GetRepos(
		context.Background(),
		&svc.RepoGetRequest{
			Id:        id,
			CreatorId: creatorId,
		},
	)
	if err != nil {
		t.Fatalf("Error for request: %v\n", err)
	}

	repos := res.Repo
	if len(repos) != 1 {
		t.Errorf("Expected to get 1 repo, Got %d repo\n", len(repos))
	}

	repo := repos[0]
	if repo.Id != id {
		t.Errorf("Expected Repo Id: %s, Got: %s\n", id, repo.Id)
	}
	if repo.Owner.Id != creatorId {
		t.Errorf("Expected Owner Id: %s, Got: %s\n", creatorId, repo.Owner.Id)
	}

	t.Logf("Reponse Repo:\n%v\n", repo)
}

func TestUserService_GetUser(t *testing.T) {
	s, l := startTestGrpcServer()
	defer s.GracefulStop()

	// grpc.WithContextDialer 매개변수로 이용하기 위한 래핑 함수
	// => net.Dial 통해 인메모리 통신
	bufconnDialer := func(ctx context.Context, addr string) (net.Conn, error) {
		return l.Dial()
	}
	client, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithInsecure(),
		grpc.WithContextDialer(bufconnDialer),
	)
	if err != nil {
		t.Fatalf("Error for creating grpc client: %v\n", err)
	}

	userClient := svc.NewUsersClient(client)

	// 인메모리 통신
	id := "mock-id"
	email := "yoonjeong@choi.com"
	resp, err := userClient.GetUser(
		context.Background(),
		&svc.UserGetRequest{
			Id:    id,
			Email: email,
		},
	)

	if err != nil {
		t.Fatalf("Error for request: %v\n", err)
	}

	respEmail := fmt.Sprintf("%s@%s", resp.User.FirstName, resp.User.LastName)
	if resp.User.Id != id ||
		respEmail != email {
		t.Errorf("Expected Id to be %s, Got: %s\n", id, resp.User.Id)
		t.Errorf("Expected Email to be %s, Got: %s\n", email, respEmail)
	}

	t.Logf("\nResponse: %v\n", resp.User)
}
