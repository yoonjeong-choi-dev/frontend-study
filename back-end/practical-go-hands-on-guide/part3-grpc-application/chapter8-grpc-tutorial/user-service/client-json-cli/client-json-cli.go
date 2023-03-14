package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"os"
	users "user-service/service"
)

func createGrpcConnection(addr string) (*grpc.ClientConn, error) {
	return grpc.DialContext(
		context.Background(),
		addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
}

func getUserServiceClient(conn *grpc.ClientConn) users.UsersClient {
	return users.NewUsersClient(conn)
}

func getUser(client users.UsersClient, u *users.UserGetRequest) (*users.UserGetResponse, error) {
	return client.GetUser(context.Background(), u)
}

func createRequestJson(id string, email string) (string, error) {
	data := map[string]string{
		"email": email,
		"id":    id,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func createUserRequestWithJsonString(jsonQuery string) (*users.UserGetRequest, error) {
	u := users.UserGetRequest{}
	input := []byte(jsonQuery)
	return &u, protojson.Unmarshal(input, &u)
}

func main() {
	var url, jsonQuery string
	var err error
	if len(os.Args) != 3 {
		fmt.Println("Not specified a gRPC server address and search query")
		fmt.Println("Default url and query would be used...")
		url = "localhost:50051"

		jsonQuery, err = createRequestJson("test", "yoonjeong@choi.com")
		if err != nil {
			log.Fatalf("Error for creating json string: %v\n", err)
		}
	} else {
		url = os.Args[1]
		jsonQuery = os.Args[2]
	}

	conn, err := createGrpcConnection(url)
	if err != nil {
		log.Fatalf("Error for grpc connection: %v\n", conn)
	}
	defer conn.Close()

	userRequest, err := createUserRequestWithJsonString(jsonQuery)
	if err != nil {
		log.Fatalf("Error for creating users.UserGetRequest")
	}

	client := getUserServiceClient(conn)
	result, err := getUser(
		client,
		userRequest,
	)

	errStatus := status.Convert(err)
	if errStatus.Code() != codes.OK {
		log.Fatalf("Request failed: %v - %v\n", errStatus.Code(), errStatus.Message())
	}

	jsonRes, err := protojson.Marshal(result)
	if err != nil {
		log.Fatalf("Error marshaling users.UserGetResponse: %v\n", err)
	}
	fmt.Printf("Response(json): %s\n", string(jsonRes))
}
