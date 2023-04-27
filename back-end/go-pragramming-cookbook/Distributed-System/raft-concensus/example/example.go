package main

import (
	"fmt"
	"net/http"
	consensus "raft-concensus"
)

func main() {
	consensus.Config(3)
	http.HandleFunc("/", consensus.StateTransitionHandler)

	fmt.Println(http.ListenAndServe(":7166", nil))
}
