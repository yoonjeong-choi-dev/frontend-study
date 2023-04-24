package gocov_testing_tool

import (
	"errors"
	"fmt"
)

type exampleStruct struct {
	Branch bool
}

func (c *exampleStruct) exampleFunc3() error {
	fmt.Println("exampleFunc3")
	if c.Branch {
		fmt.Println("branch was set")
		return errors.New("bad branch")
	}
	return nil
}
