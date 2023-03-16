package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"io"
	"log"
	"net"
	svc "streaming/service"
	"streaming/utils"
	"testing"
)

func startTestGrpcServer() *bufconn.Listener {
	l := bufconn.Listen(10)
	s := grpc.NewServer()

	registerService(s)
	go func() {
		defer s.GracefulStop()
		log.Fatal(s.Serve(l))
	}()
	return l
}

func createRepoClient(l *bufconn.Listener) (svc.ReposClient, error) {
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

	return svc.NewReposClient(client), nil
}

func TestRepoService_GetRepos(t *testing.T) {
	l := startTestGrpcServer()

	client, err := createRepoClient(l)
	if err != nil {
		t.Fatalf("Error for creating dial: %v\n", err)
	}

	// Testing
	id := "id"
	creatorId := "creator"
	stream, err := client.GetRepos(
		context.Background(),
		&svc.RepoGetRequest{Id: id, CreatorId: creatorId},
	)

	if err != nil {
		t.Fatalf("Expected no err but got %v\n", err)
	}

	var repos []*svc.Repository
	for {
		repo, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Error for reading stream: %v\n", err)
		}
		repos = append(repos, repo.Repo)
	}

	if len(repos) != 5 {
		t.Fatalf("Expected to get 5 repos but got %d repos\n", len(repos))
	}

	for idx, repo := range repos {
		if id != repo.Id {
			t.Errorf("Expected Id: %s, but got %s\n", id, repo.Id)
		}

		gotRepoName := repo.Name
		expected := fmt.Sprintf("repo-name-%d", idx+1)

		if gotRepoName != expected {
			t.Errorf("Expected Name: %s, but got %s\n", expected, gotRepoName)
		}
	}

	t.Logf("%s\n", utils.GetJsonStringUnsafe(repos))
}

func TestRepoService_CreateBuild(t *testing.T) {
	l := startTestGrpcServer()
	client, err := createRepoClient(l)
	if err != nil {
		t.Fatalf("Error for creating dial: %v\n", err)
	}

	repoName := "test-repo"
	request := svc.Repository{
		Id:   "id",
		Name: repoName,
	}
	stream, err := client.CreateBuild(context.Background(), &request)
	if err != nil {
		t.Fatalf("Expected nil err but got %v\n", err)
	}

	var buildLogs []*svc.RepoBuildLog
	for {
		buildLog, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fatalf("Expected nil err but got %v\n", err)
		}
		buildLogs = append(buildLogs, buildLog)
	}

	if len(buildLogs) != 7 {
		t.Fatalf("Expected to get 7 repos but got %d repos\n", len(buildLogs))
	}

	var expectedLog string
	for idx, buildLog := range buildLogs {
		if idx == 0 {
			expectedLog = fmt.Sprintf("Starting Building For %s", repoName)
		} else if idx == len(buildLogs)-1 {
			expectedLog = fmt.Sprintf("Finished Building For %s", repoName)
		} else {
			expectedLog = fmt.Sprintf("BuildLogLine-%d", idx)
		}

		if buildLog.Log != expectedLog {
			t.Errorf("Expected: %s, Got: %s\n", expectedLog, buildLog.Log)
		}

		if err := buildLog.Timestamp.CheckValid(); err != nil {
			t.Errorf("Invalid timestamp: %v\n", err)
		}
	}

	t.Logf("%s\n", utils.GetJsonStringUnsafe(buildLogs))
}
