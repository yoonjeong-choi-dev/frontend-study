package handlers

import (
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"grpc-json-api/keyvalue"

	"github.com/apex/log"
	"net/http"
)

func (c *Controller) GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	kvReq := keyvalue.GetKeyValueRequest{Key: key}

	kvRes, err := c.Get(r.Context(), &kvReq)
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

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
