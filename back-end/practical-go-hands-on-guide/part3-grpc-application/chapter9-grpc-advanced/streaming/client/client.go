package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	svc "streaming/service"
	"streaming/utils"
)

func createGrpcConnection(addr string) (*grpc.ClientConn, error) {
	return grpc.DialContext(
		context.Background(),
		addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
}

func getUserServiceClient(conn *grpc.ClientConn) svc.UsersClient {
	return svc.NewUsersClient(conn)
}

func getRepoServiceClient(conn *grpc.ClientConn) svc.ReposClient {
	return svc.NewReposClient(conn)
}

// Type 1. Unary
func getUser(client svc.UsersClient, req *svc.UserGetRequest) (*svc.UserGetResponse, error) {
	return client.GetUser(context.Background(), req)
}

// Type 2. Server Side Stream
func getRepo(client svc.ReposClient, req *svc.RepoGetRequest) ([]*svc.Repository, error) {
	// stream: go grpc 플러그인이 구현한 Repos_GetReposClient 인터페이스 구현체
	stream, err := client.GetRepos(context.Background(), req)
	if err != nil {
		return nil, err
	}

	var repos []*svc.Repository
	for {
		repo, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error reading stream: %v\n", err)
			return repos, err
		}

		repos = append(repos, repo.Repo)
	}
	return repos, nil
}

func createBuild(client svc.ReposClient, req *svc.Repository) ([]*svc.RepoBuildLog, error) {
	stream, err := client.CreateBuild(context.Background(), req)
	if err != nil {
		return nil, err
	}

	var logs []*svc.RepoBuildLog
	for {
		log, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		logs = append(logs, log)
	}
	return logs, nil
}

// Type 3. Client Side Streaming
func createRepos(client svc.ReposClient, req []*svc.RepoCreateRequest) (*svc.RepoCreateResponse, error) {
	stream, err := client.CreateRepos(context.Background())
	if err != nil {
		return nil, err
	}

	for _, r := range req {
		err := stream.Send(r)
		if err != nil {
			return nil, err
		}
	}
	return stream.CloseAndRecv()
}

func main() {
	conn, err := createGrpcConnection(":50051")
	if err != nil {
		log.Fatalf("Error for grpc connection: %v\n", err)
	}

	fmt.Println("Example for User Service")
	userClient := getUserServiceClient(conn)
	result, err := getUser(userClient, &svc.UserGetRequest{Email: "yoonjeong@choi"})
	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		fmt.Printf("Request Failed: %v\n", utils.GetJsonStringUnsafe(errStatus))
	} else {
		fmt.Printf("Response: %s\n", utils.GetJsonStringUnsafe(result))
	}

	fmt.Println("Example for Repo Service - Get Repo")
	repoClient := getRepoServiceClient(conn)
	repoResult, err := getRepo(repoClient, &svc.RepoGetRequest{Id: "yj", CreatorId: "yjchoi7166"})
	errStatus = status.Convert(err)
	if errStatus.Code() != codes.OK {
		fmt.Printf("Request Failed: %v\n", utils.GetJsonStringUnsafe(errStatus))
	} else {
		fmt.Printf("Response: %d data\n", len(repoResult))
		fmt.Printf("Response: %s\n", utils.GetJsonStringUnsafe(repoResult))
	}

	fmt.Println("Example for Repo Service - Create Log")
	logResult, err := createBuild(repoClient, &svc.Repository{
		Id:   "Build Id",
		Name: "Anonymous",
	})
	errStatus = status.Convert(err)
	if errStatus.Code() != codes.OK {
		fmt.Printf("Request Failed: %v\n", utils.GetJsonStringUnsafe(errStatus))
	} else {
		fmt.Printf("Response: %d data\n", len(logResult))
		fmt.Printf("Response: %s\n", utils.GetJsonStringUnsafe(logResult))
	}

	fmt.Println("Example for Repo Service - Create Repo")
	//var createRepoReq =
	createRes, err := createRepos(repoClient, []*svc.RepoCreateRequest{
		{Id: "yj", Name: "moloco"},
		{Id: "123", Name: "gangnam"},
		{Id: "choi", Name: "test-repo"},
	})
	errStatus = status.Convert(err)
	if errStatus.Code() != codes.OK {
		fmt.Printf("Request Failed: %v\n", utils.GetJsonStringUnsafe(errStatus))
	} else {
		createResult := createRes.Repos
		fmt.Printf("Response: %d data\n", len(createResult))
		fmt.Printf("Response: %s\n", utils.GetJsonStringUnsafe(createResult))
	}

}
