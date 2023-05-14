package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	res, err := http.Head("https://www.time.gov/")
	if err != nil {
		panic(err)
	}
	defer func() { _ = res.Body.Close() }()

	now := time.Now().Round(time.Second)
	// Date Header: 서버가 응답을 생성한 시간
	dateHeader := res.Header.Get("Date")
	if dateHeader == "" {
		fmt.Println("No Date Header from time.gov")
		return
	}

	date, err := time.Parse(time.RFC1123, dateHeader)
	if err != nil {
		panic(err)
	}

	fmt.Printf("time.gov: %s (skew %s)\n", date, now.Sub(date))
}
