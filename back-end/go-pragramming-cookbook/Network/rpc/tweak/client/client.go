package main

import (
	"fmt"
	"log"
	"net/rpc"
	"rpc/tweak"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:7166")
	if err != nil {
		log.Fatalln("fail to connect:", err)
	}

	var result string
	err = client.Call("StringTweaker.Echo", "Echo Test", &result)
	if err != nil {
		log.Fatalln("client call with error:", err)
	}
	fmt.Printf("Received: %s\n", result)

	args := tweak.Args{
		String:  "this string should be uppercase",
		ToUpper: true,
		Reverse: false,
	}

	err = client.Call("StringTweaker.Tweak", args, &result)
	if err != nil {
		log.Fatalln("client call with error:", err)
	}
	fmt.Printf("Received: %s\n", result)

	another := tweak.Args{
		String:  "reverse test",
		ToUpper: false,
		Reverse: true,
	}
	err = client.Call("StringTweaker.Tweak", another, &result)
	if err != nil {
		log.Fatalln("client call with error:", err)
	}
	fmt.Printf("Received: %s\n", result)
}
