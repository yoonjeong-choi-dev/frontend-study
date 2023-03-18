package main

import (
	svc "binary-data/service"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

type repoService struct {
	svc.UnimplementedRepoServer
}

func (s *repoService) Create(stream svc.Repo_CreateServer) error {
	var repoContext *svc.RepoContext
	var data []byte
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Error(codes.Unknown, err.Error())
		}

		switch t := req.Body.(type) {
		// oneof 키워드 내부의 필드들은 언더바로 생성됨
		case *svc.RepoCreateRequest_Context:
			repoContext = req.GetContext()
		case *svc.RepoCreateRequest_Data:
			curData := req.GetData()
			data = append(data, curData...)
		case nil:
			return status.Error(
				codes.InvalidArgument,
				"message must contain one of context or data",
			)
		default:
			return status.Errorf(
				codes.FailedPrecondition,
				"unexpected message type: %s",
				t,
			)
		}
	}

	if repoContext.CreatorName == "" || repoContext.FileName == "" {
		return status.Error(
			codes.InvalidArgument,
			"file name and creator name required",
		)
	}

	path := fmt.Sprintf("tmp/%s", repoContext.CreatorName)

	fileSize, err := saveFile(path, repoContext.FileName, data)
	if err != nil {
		return status.Error(
			codes.Internal,
			err.Error(),
		)
	}

	repo := svc.Repository{
		Id:   fmt.Sprintf("%s-%s", repoContext.CreatorName, repoContext.FileName),
		Name: repoContext.FileName,
		Url:  fmt.Sprintf("%s/%s", path, repoContext.FileName),
	}
	res := svc.RepoCreateResponse{
		Repo: &repo,
		Size: int32(fileSize),
	}
	return stream.SendAndClose(&res)
}

func saveFile(path string, fileName string, data []byte) (int, error) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return -1, err
	}

	file, err := os.Create(filepath.Join(path, fileName))
	if err != nil {
		return -1, err
	}
	defer file.Close()

	return file.Write(data)
}

func registerService(s *grpc.Server) {
	svc.RegisterRepoServer(s, &repoService{})
}

func startServer(s *grpc.Server, l net.Listener) error {
	return s.Serve(l)
}

func main() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error for TCP Litening: %v\n", err)
	}

	server := grpc.NewServer()
	registerService(server)
	log.Fatal(startServer(server, l))
}
