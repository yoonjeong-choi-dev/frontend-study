package main

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func GuessHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error reading request body"))
		return
	}

	msg := r.FormValue("message")

	// hashed for "password"
	real := []byte("$2a$10$2ovnPWuIjMx2S0HvCxP/mutzdsGhyt8rq/JqnJg/6OyC3B0APMGlK")

	if err := bcrypt.CompareHashAndPassword(real, []byte(msg)); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong message. try again"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("right message"))
}
