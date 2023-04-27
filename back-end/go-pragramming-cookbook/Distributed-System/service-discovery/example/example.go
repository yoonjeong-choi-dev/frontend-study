package main

import (
	"github.com/hashicorp/consul/api"
	discovery "service-discovery"
)

func main() {
	config := api.DefaultConfig()

	// run-consul.sh 실행을 통해 실행되는 서버
	config.Address = "localhost:8500"

	client, err := discovery.NewClient(config, "localhost", "yj-service", 7166)
	if err != nil {
		panic(err)
	}

	if err := discovery.ExampleWithClient(client); err != nil {
		panic(err)
	}
}
