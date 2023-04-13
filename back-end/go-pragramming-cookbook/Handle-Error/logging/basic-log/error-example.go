package main

import (
	"github.com/pkg/errors"
	"log"
)

func OriginalError() error {
	return errors.New("initial error occurred")
}

func PassThroughError() error {
	err := OriginalError()
	return errors.Wrap(err, "in passThroughError")
}

func FinalDestination() {
	err := PassThroughError()
	if err != nil {
		log.Printf("an error occured: %s\n", err.Error())
	}
}
