package main

import (
	"context"
	"fmt"
	"io"
	svc "robust-app/service"
)

func (s *userService) RisePanicTest() {
	fmt.Fprintln(s.writer, "<RisePanicTest>")

	req := &svc.PanicTestMessage{
		Message: "Test",
	}

	res, err := s.client.RisePanic(context.Background(), req)
	printResponse(s.writer, res, err)

	req.Message = "panic"
	res, err = s.client.RisePanic(context.Background(), req)
	printResponse(s.writer, res, err)
}

func (s *userService) RiseServerStreamPanic() {
	fmt.Fprintln(s.writer, "<RiseServerStreamPanic>")

	req := &svc.PanicTestMessage{
		Message: "Test",
	}
	stream, err := s.client.RiseStreamPanic(context.Background(), req)
	if err != nil {
		printResponse(s.writer, nil, err)
		return
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		printResponse(s.writer, res, err)
		if err != nil {
			return
		}
	}
}
