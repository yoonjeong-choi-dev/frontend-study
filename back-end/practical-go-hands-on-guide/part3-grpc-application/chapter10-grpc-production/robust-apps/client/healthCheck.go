package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	healthsvc "google.golang.org/grpc/health/grpc_health_v1"
	"io"
	"time"
)

type healthService struct {
	client healthsvc.HealthClient
	conn   *grpc.ClientConn
	cancel context.CancelFunc
	reader io.Reader
	writer io.Writer
}

func (s *healthService) InitClient(addr string) error {
	if s.writer != nil {
		fmt.Fprintf(s.writer, "[HealthService]Connecting to server on %s\n", addr)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	s.cancel = cancel

	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithInsecure(),
		// 일시적인 네트워크 문제이 대해서 무한정 대기
		grpc.WithBlock(),
		// 영구적인 네트워크 에러에 대해서 연결 재시도 X
		grpc.FailOnNonTempDialError(true),
		// 연결 전에 일시적인 에러가 발생하고 컨텍스트 만료(여기서는 타임아웃) 시, 해당 에러 반환
		grpc.WithReturnConnectionError(),
	)

	if err != nil {
		return err
	}

	s.conn = conn
	s.client = healthsvc.NewHealthClient(conn)
	return nil
}

func (s *healthService) InitInteraction(r io.Reader, w io.Writer) {
	s.reader = r
	s.writer = w
}

func (s *healthService) CheckUserService() (*healthsvc.HealthCheckResponse, error) {
	return s.client.Check(
		context.Background(),
		&healthsvc.HealthCheckRequest{
			Service: "Users",
		},
	)
}

func (s *healthService) GetUserServiceStatus() bool {
	res, err := s.CheckUserService()
	fmt.Fprintf(s.writer, "Health Check For User Serivce")
	printResponse(s.writer, res.Status.String(), err)

	return err != nil || res.Status != healthsvc.HealthCheckResponse_SERVING
}
