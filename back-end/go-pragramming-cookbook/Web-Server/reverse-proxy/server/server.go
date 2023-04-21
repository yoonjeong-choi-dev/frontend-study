package main

import (
	"fmt"
	"net/http"
	reverse_proxy "reverse-proxy"
)

func main() {
	proxy := &reverse_proxy.Proxy{
		Client:    http.DefaultClient,
		ServerURL: "https://www.google.com",
	}

	http.Handle("/", proxy)
	fmt.Println(http.ListenAndServe(":7166", nil))
}
