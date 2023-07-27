package main

import (
	"context"
	"fmt"
	"housework-app/housework/v1"
	"strconv"
	"strings"
)

type MaidService struct {
	client housework.RobotMaidClient
}

func (m *MaidService) list(ctx context.Context) error {
	chores, err := m.client.List(ctx, new(housework.Empty))
	if err != nil {
		return err
	}

	if len(chores.Chores) == 0 {
		fmt.Println("You have nothing to do!")
		return nil
	}

	fmt.Println("#\t[X]\tDescription")
	for i, chore := range chores.Chores {
		complete := ""
		if chore.Complete {
			complete = "X"
		}
		fmt.Printf("%d\t[%s]\t%s\n", i+1, complete, chore.Description)
	}

	return nil
}

func (m *MaidService) add(ctx context.Context, s string) error {
	chores := new(housework.Chores)

	for _, chore := range strings.Split(s, ",") {
		if desc := strings.TrimSpace(chore); desc != "" {
			chores.Chores = append(chores.Chores, &housework.Chore{Description: desc})
		}
	}

	var err error
	if len(chores.Chores) > 0 {
		_, err = m.client.Add(ctx, chores)
	}
	return err
}

func (m *MaidService) complete(ctx context.Context, s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	req := &housework.CompleteRequest{ChoreNumber: int32(i)}

	_, err = m.client.Complete(ctx, req)
	return err
}
