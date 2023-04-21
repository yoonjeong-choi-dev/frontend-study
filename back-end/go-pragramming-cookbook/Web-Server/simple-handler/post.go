package simple_handler

import (
	"encoding/json"
	"net/http"
)

type GreetingResponse struct {
	Payload struct {
		Greeting string `json:"greeting,omitempty"`
		Name     string `json:"name,omitempty"`
		Error    string `json:"error,omitempty"`
	} `json:"payload"`
	Successful bool `json:"successful"`
}

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res GreetingResponse
	if err := r.ParseForm(); err != nil {
		res.Payload.Error = "bad request"
		res.Successful = false
		if payload, err := json.Marshal(res); err == nil {
			w.Write(payload)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	name := r.FormValue("name")
	greeting := r.FormValue("greeting")

	w.WriteHeader(http.StatusOK)
	res.Payload.Name = name
	res.Payload.Greeting = greeting
	res.Successful = true

	if payload, err := json.Marshal(res); err == nil {
		w.Write(payload)
	}
}
