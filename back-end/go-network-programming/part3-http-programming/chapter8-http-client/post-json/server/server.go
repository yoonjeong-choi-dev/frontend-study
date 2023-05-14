package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type user struct {
	Name string
	Age  int64
}

func postUserHandler(w http.ResponseWriter, r *http.Request) {
	defer func(r io.ReadCloser) {
		_, _ = io.Copy(ioutil.Discard, r)
		_ = r.Close()
	}(r.Body)

	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	var u user
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "error for decoding body", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(fmt.Sprintf("Your name is %s, and age is %d", u.Name, u.Age)))
}

func main() {
	http.HandleFunc("/", postUserHandler)
	fmt.Println(http.ListenAndServe(":7166", nil))
}
