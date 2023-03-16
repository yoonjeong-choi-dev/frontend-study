package main

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	svc "streaming/service"
	"streaming/utils"
	"time"
)

type repoService struct {
	svc.UnimplementedReposServer
}

// GetRepos Server Side Streaming
func (s *repoService) GetRepos(
	in *svc.RepoGetRequest,
	stream svc.Repos_GetReposServer,
) error {
	log.Printf("[Streaming] Recieved request for repo with Id: %s, CreatedId: %s\n",
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

	count := 1
	for {
		repo.Name = fmt.Sprintf("repo-name-%d", count)
		repo.Url = fmt.Sprintf("github.com/yoonjeong-choi-dev/stream/%s", repo.Name)

		res := svc.RepoGetResponse{Repo: &repo}

		// stream.Send : GO grpc 플러그인이 구현한 Repos_GetReposServer 인터페이스의 구현체
		if err := stream.Send(&res); err != nil {
			return err
		}
		if count >= 5 {
			break
		}
		count++
	}
	return nil
}

// CreateBuild Server Side Streaming
func (s *repoService) CreateBuild(
	in *svc.Repository,
	stream svc.Repos_CreateBuildServer,
) error {
	log.Printf("[Streaming] Recieved Repository Request with id: %s\n", in.Id)

	buildLog := svc.RepoBuildLog{
		Log:       fmt.Sprintf("Starting Building For %s", in.Name),
		Timestamp: timestamppb.Now(),
	}
	if err := stream.Send(&buildLog); err != nil {
		return err
	}

	count := 1
	for {
		time.Sleep(500 * time.Millisecond)
		buildLog.Log = fmt.Sprintf("BuildLogLine-%d", count)
		buildLog.Timestamp = timestamppb.Now()

		if err := stream.Send(&buildLog); err != nil {
			return err
		}
		if count >= 5 {
			break
		}
		count++
	}

	buildLog = svc.RepoBuildLog{
		Log:       fmt.Sprintf("Finished Building For %s", in.Name),
		Timestamp: timestamppb.Now(),
	}
	if err := stream.Send(&buildLog); err != nil {
		return err
	}
	return nil
}

// CreateRepos Client Side Stream
func (s *repoService) CreateRepos(stream svc.Repos_CreateReposServer) error {
	log.Println("Start Listening Client Streaming Data")
	var res []*svc.Repository
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Error for reading streaming: %s\n", utils.GetJsonStringUnsafe(err))
			return status.Error(
				codes.Unknown,
				err.Error())
		}

		repo := &svc.Repository{
			Id:   fmt.Sprintf("%s-%d", req.Id, count),
			Name: fmt.Sprintf("%s-%d", req.Name, count),
			Url:  fmt.Sprintf("github.com/yoonjeong-choi-dev/%s/%d", req.Name, count),
			Owner: &svc.User{
				Id:        req.Id,
				FirstName: "Yoonjeong",
				LastName:  "Choi",
				Age:       31,
			},
		}
		res = append(res, repo)
		count++
	}

	log.Println("Finished Listening Client Streaming Data")
	return stream.SendAndClose(&svc.RepoCreateResponse{
		Repos: res,
	})
}
