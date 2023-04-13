package main

import (
	"errors"
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"os"
)

type CustomHandler struct {
	Id      string
	handler log.Handler
}

func (h *CustomHandler) HandleLog(e *log.Entry) error {
	e.WithField("id", h.Id)
	return h.handler.HandleLog(e)
}

func throwAndLoggingError() error {
	err := errors.New("some error occurred")
	log.WithField("id", "some-error-id").Trace("ThrowError").Stop(&err)
	return err
}

func ApexExample() {
	log.SetHandler(&CustomHandler{"yj-handler", text.New(os.Stdout)})

	err := throwAndLoggingError()

	log.WithError(err).Error("an error occurred")
}
