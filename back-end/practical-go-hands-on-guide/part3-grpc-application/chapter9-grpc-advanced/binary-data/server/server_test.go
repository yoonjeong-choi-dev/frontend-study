package main

import (
	svc "binary-data/service"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"io"
	"log"
	"net"
	"strings"
	"testing"
	"time"
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

func TestRepoService_Create(t *testing.T) {
	l := startTestGrpcServer()

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
		log.Fatalf("Error for creating grpc client: %v\n", err)
	}

	repoClient := svc.NewRepoClient(client)

	stream, err := repoClient.Create(context.Background())
	if err != nil {
		log.Fatalf("Error for creating stream: %v\n", err)
	}

	repoContext := svc.RepoCreateRequest_Context{
		Context: &svc.RepoContext{
			CreatorName: "yoonjeong",
			FileName:    fmt.Sprintf("test-repo-%d.txt", time.Now().Unix()),
		},
	}

	req := svc.RepoCreateRequest{Body: &repoContext}
	err = stream.Send(&req)
	if err != nil {
		log.Fatalf("Error for sending Repo Context: %v\n", err)
	}

	testData := "This is Test Data"
	dataStream := strings.NewReader(testData)
	for {
		b, err := dataStream.ReadByte()
		if err == io.EOF {
			break
		}

		reqData := svc.RepoCreateRequest_Data{Data: []byte{b}}
		req := svc.RepoCreateRequest{Body: &reqData}

		err = stream.Send(&req)
		if err != nil {
			log.Fatalf("Error for sending binary data: %v\n", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Error for closing connection: %v\n", err)
	}

	expectedSize := int32(len(testData))
	if expectedSize != res.Size {
		t.Errorf("Expected size: %d but got: %d\n", expectedSize, res.Size)
	}

	t.Logf("Response: %s\n", getJsonStringUnsafe(res))
}

func getJsonStringUnsafe(v interface{}) string {
	result, _ := json.MarshalIndent(v, "", " ")
	return string(result)
}
