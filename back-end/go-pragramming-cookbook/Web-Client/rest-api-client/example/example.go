package main

import "restapi"

func main() {
	if err := restapi.ExecExample(); err != nil {
		panic(err)
	}
}
