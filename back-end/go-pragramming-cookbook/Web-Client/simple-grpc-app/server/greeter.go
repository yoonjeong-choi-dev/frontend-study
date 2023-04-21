package main

import (
	"context"
	"fmt"
	"greeter"
)

type Greeter struct {
	Exclaim bool
	greeter.UnimplementedGreeterServiceServer
}

func (g *Greeter) Greet(ctx context.Context, r *greeter.GreetRequest) (*greeter.GreetResponse, error) {
	var exclaim string
	if g.Exclaim {
		exclaim = "!"
	} else {
		exclaim = "."
	}
	msg := fmt.Sprintf("%s, %s%s", r.GetGreeting(), r.GetName(), exclaim)

	return &greeter.GreetResponse{Response: msg}, nil
}
