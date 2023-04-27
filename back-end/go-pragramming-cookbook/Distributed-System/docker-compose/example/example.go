package main

import docker_compose "docker-compose"

func main() {
	if err := docker_compose.ExecExample("mongodb://mongodb:27017"); err != nil {
		panic(err)
	}
}
