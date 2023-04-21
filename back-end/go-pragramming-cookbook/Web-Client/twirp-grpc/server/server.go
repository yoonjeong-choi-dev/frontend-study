package main

import (
	"fmt"
	"net/http"
	"service"
)

func main() {
	server := &Greeter{Exclaim: false}
	twirpHandler := service.NewGreeterServiceServer(server, nil)

	fmt.Printf("Listen on port 7166 with twirp!")
	fmt.Println(http.ListenAndServe(":7166", twirpHandler))
}
