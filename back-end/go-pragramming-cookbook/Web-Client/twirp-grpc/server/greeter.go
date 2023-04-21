package main

import (
	"context"
	"fmt"
	"service"
)

type Greeter struct {
	Exclaim bool
}

func (g *Greeter) Greet(ctx context.Context, r *service.GreetRequest) (*service.GreetResponse, error) {
	var exclaim string
	if g.Exclaim {
		exclaim = "!"
	} else {
		exclaim = "."
	}
	msg := fmt.Sprintf("%s, %s%s", r.GetGreeting(), r.GetName(), exclaim)

	return &service.GreetResponse{Response: msg}, nil
}
