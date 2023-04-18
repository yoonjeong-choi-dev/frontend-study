package client

import (
	"fmt"
	"net/http"
)

type Controller struct {
	*http.Client
}

func (c *Controller) DoOperations(url string) error {
	resp, err := c.Client.Get(url)
	if err != nil {
		return err
	}

	fmt.Printf("Controller.DoOperations Result Code from %s: %d\n", url, resp.StatusCode)
	return nil
}
