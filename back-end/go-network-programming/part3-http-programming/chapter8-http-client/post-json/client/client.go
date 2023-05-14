package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type user struct {
	Name string
	Age  int64
}

const host = "http://localhost:7166"

func main() {
	// GET
	res, err := http.Get(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get Method Response Status: %s\n", res.Status)

	// POST
	buf := new(bytes.Buffer)
	u := user{Name: "Yoonjeong", Age: 31}
	if err := json.NewEncoder(buf).Encode(&u); err != nil {
		panic(err)
	}

	res, err = http.Post(host, "application/json", buf)
	if err != nil {
		panic(err)
	}
	defer func() { _ = res.Body.Close() }()

	fmt.Printf("Post Method Response Status: %s\n", res.Status)
	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Post Method Response Body: %s\n", string(data))
}
