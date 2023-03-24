package main

import (
	"context"
	"fmt"
	svc "robust-app/service"
)

// RisePanic Unary Pattern
func (s *userService) RisePanic(
	ctx context.Context,
	in *svc.PanicTestMessage,
) (*svc.PanicTestMessage, error) {
	if in.Message == "panic" {
		panic("Panic For Unary Pattern")
	}

	res := svc.PanicTestMessage{
		Message: fmt.Sprintf("Echo: %s", in.Message),
	}
	return &res, nil
}

// RiseStreamPanic Server-Side Stream
func (s *userService) RiseStreamPanic(
	in *svc.PanicTestMessage,
	stream svc.Users_RiseStreamPanicServer,
) error {
	res := svc.PanicTestMessage{}

	step := 1
	for {
		res.Message = fmt.Sprintf("Echo - %d: %s", step, in.Message)
		if err := stream.Send(&res); err != nil {
			return err
		}

		if step == 3 {
			panic("Panic For Server Side Streaming")
		}

		step++
	}

	return nil
}
