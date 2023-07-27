package main

import (
	"context"
	"fmt"
	"housework-app/housework/v1"
	"sync"
)

type RobotMaid struct {
	mu     sync.Mutex
	chores []*housework.Chore
	housework.UnimplementedRobotMaidServer
}

func (r *RobotMaid) Add(_ context.Context, chores *housework.Chores) (
	*housework.Response, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.chores = append(r.chores, chores.Chores...)
	return &housework.Response{Message: "ok"}, nil
}

func (r *RobotMaid) Complete(_ context.Context, in *housework.CompleteRequest) (
	*housework.Response, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.chores == nil || in.ChoreNumber < 1 || int(in.ChoreNumber) > len(r.chores) {
		return nil, fmt.Errorf("chore %d not found", in.ChoreNumber)
	}

	r.chores[in.ChoreNumber-1].Complete = true
	return &housework.Response{Message: "ok"}, nil
}

func (r *RobotMaid) List(_ context.Context, _ *housework.Empty) (
	*housework.Chores, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.chores == nil {
		r.chores = make([]*housework.Chore, 0)
	}

	return &housework.Chores{Chores: r.chores}, nil
}
