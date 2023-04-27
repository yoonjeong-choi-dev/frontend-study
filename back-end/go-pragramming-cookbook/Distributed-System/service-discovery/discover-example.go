package service_discovery

import (
	"encoding/json"
	"fmt"
)

func ExampleWithClient(client Client) error {
	if err := client.Register([]string{"Go", "yoonjeong", "discovery", "yj-tag"}); err != nil {
		return err
	}

	serviceEntries, _, err := client.Service("yj-discovery", "yoonjeong")
	if err != nil {
		return err
	}

	for _, entry := range serviceEntries {
		fmt.Printf("Service: %s\n", getJsonString(entry.Service))
	}
	return nil
}

func getJsonString(v interface{}) string {
	val, _ := json.MarshalIndent(v, "", " ")
	return string(val)
}
