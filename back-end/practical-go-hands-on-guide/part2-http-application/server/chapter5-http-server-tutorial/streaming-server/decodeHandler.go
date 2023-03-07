package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type eventLog struct {
	UserIP string `json:"user_ip"`
	Event  string `json:"event"`
}

func decodeLogStreamHandler(w http.ResponseWriter, r *http.Request) {
	var jsonParseErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Exercise 5.2

	for {
		var logData eventLog
		err := decoder.Decode(&logData)

		if err == io.EOF {
			break
		}

		// parsing error
		if errors.As(err, &jsonParseErr) {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			continue
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			continue
		}
		fmt.Printf("[log]%s - %s\n", logData.UserIP, logData.Event)
		fmt.Fprintln(w, "OK")
	}
	fmt.Fprintln(w, "Streaming Processing Done")
}
