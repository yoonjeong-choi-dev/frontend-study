package handlers

import (
	"encoding/json"
	"github.com/apex/log"
	"grpc-json-api/keyvalue"
	"net/http"
)

func (c *Controller) SetHandler(w http.ResponseWriter, r *http.Request) {
	var kvReq keyvalue.SetKeyValueRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&kvReq); err != nil {
		log.Errorf("failed to decode: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	kvRes, err := c.Set(r.Context(), &kvReq)
	if err != nil {
		log.Errorf("failed to call grpc: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(kvRes)
	if err != nil {
		log.Errorf("failed to marshal: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
