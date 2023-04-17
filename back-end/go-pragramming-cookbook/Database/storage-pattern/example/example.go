package main

import "storage"

func main() {
	if err := storage.MongoExecExample(); err != nil {
		panic(err)
	}
}
