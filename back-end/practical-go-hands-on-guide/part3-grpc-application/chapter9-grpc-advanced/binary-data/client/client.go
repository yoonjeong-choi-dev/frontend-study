package main

import (
	svc "binary-data/service"
	"context"
	"errors"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"path/filepath"
)

type appConfig struct {
	filePath   string
	serverUrl  string
	clientName string
	fileName   string
}

func setupFlag(w io.Writer, args []string) (appConfig, error) {
	config := appConfig{}

	fs := flag.NewFlagSet("grpc-cli", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&config.filePath, "file-path", "", "File Path to upalod")
	fs.StringVar(&config.clientName, "user-name", "anonymous", "User name")
	fs.StringVar(&config.fileName, "file-name", "", "File name to save")

	err := fs.Parse(args)
	if err != nil {
		return config, err
	}

	if len(config.filePath) == 0 {
		return config, errors.New("file path is empty")
	}

	if len(config.fileName) == 0 {
		config.fileName = filepath.Base(config.filePath)
	}

	if fs.NArg() != 1 {
		fs.Usage()
		return config, errors.New("must specify server URL as positional argument only")
	}

	config.serverUrl = fs.Arg(0)

	log.Printf("App Config: %#v\n", config)
	return config, nil
}

func createGrpcConnection(addr string) (*grpc.ClientConn, error) {
	return grpc.DialContext(
		context.Background(),
		addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
}

func getRepoServiceClient(conn *grpc.ClientConn) svc.RepoClient {
	return svc.NewRepoClient(conn)
}

func uploadFile(config appConfig, client svc.RepoClient) (*svc.RepoCreateResponse, error) {
	stream, err := client.Create(context.Background())
	if err != nil {
		return nil, err
	}

	// Step 1. Send meta data
	repoContext := svc.RepoCreateRequest_Context{
		Context: &svc.RepoContext{
			CreatorName: config.clientName,
			FileName:    config.fileName,
		},
	}
	req := svc.RepoCreateRequest{Body: &repoContext}
	err = stream.Send(&req)
	if err != nil {
		return nil, err
	}

	// Step 2. Send binary stream data
	file, err := os.Open(config.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	size := 32 * 1024
	streamBuf := make([]byte, size)
	for {
		nBytes, err := file.Read(streamBuf)
		if err == io.EOF {
			break
		}

		reqData := svc.RepoCreateRequest_Data{
			Data: streamBuf[:nBytes],
		}
		req := svc.RepoCreateRequest{Body: &reqData}
		err = stream.Send(&req)
		if err != nil {
			return nil, err
		}
	}

	return stream.CloseAndRecv()
}

func main() {
	config, err := setupFlag(os.Stdout, os.Args[1:])
	if err != nil {
		log.Fatalf("Error for args: %v\n", err)
	}

	conn, err := createGrpcConnection(config.serverUrl)
	if err != nil {
		log.Fatalf("Error for creating grpc conn :%v\n", err)
	}
	defer conn.Close()

	client := getRepoServiceClient(conn)
	res, err := uploadFile(config, client)
	if err != nil {
		log.Fatalf("Error for uploading: %v\n", err)
	}

	fmt.Printf("Upload Suceess: %v\n", res)
}
