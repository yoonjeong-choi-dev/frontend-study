package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	svc "multiple-services/service"
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

func getUser(client svc.UsersClient, req *svc.UserGetRequest) (*svc.UserGetResponse, error) {
	return client.GetUser(context.Background(), req)
}

func getRepoServiceClient(conn *grpc.ClientConn) svc.RepoClient {
	return svc.NewRepoClient(conn)
}

func getRepo(client svc.RepoClient, req *svc.RepoGetRequest) (*svc.RepoGetResponse, error) {
	return client.GetRepos(context.Background(), req)
}

func main() {
	conn, err := createGrpcConnection(":50051")
	if err != nil {
		log.Fatal("Error for grpc connection: %v\n", err)
	}

	exampleGetUser(conn)
	exampleGetRepo(conn)
}

func getJsonString(v interface{}) string {
	vJson, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		fmt.Printf("Marshal Error: %s\n", err.Error())
		return ""
	}
	return string(vJson)
}

func exampleGetUser(conn *grpc.ClientConn) {
	fmt.Println("svc.GetUser Examples")
	emails := []string{
		"yoonjeong@choi",
		"invalid-data",
		"",
	}
	client := getUserServiceClient(conn)

	for idx, email := range emails {
		fmt.Printf("Example %d - email: %s\n", idx, email)
		result, err := getUser(
			client,
			&svc.UserGetRequest{Email: email},
		)

		errStatus := status.Convert(err)
		if errStatus.Code() == codes.OK {
			fmt.Printf("Response:\n%s\n", getJsonString(result.User))
		} else {
			fmt.Printf("Request failed: %v - %v\n", errStatus.Code(), errStatus.Message())
		}
	}
}

func exampleGetRepo(conn *grpc.ClientConn) {
	fmt.Println("svc.GetRepos Examples")

	client := getRepoServiceClient(conn)
	result, err := getRepo(
		client,
		&svc.RepoGetRequest{
			Id:        "repo-id",
			CreatorId: "yoonjeong-choi-dev",
		},
	)

	errStatus := status.Convert(err)
	if errStatus.Code() == codes.OK {
		fmt.Printf("Response:\n%s\n", getJsonString(result))
	} else {
		fmt.Printf("Request Failed: %v - %v\n", errStatus.Code(), errStatus.Message())
	}
}
