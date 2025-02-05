package docker

import (
	"encoding/json"
	"net/http"
	"time"
)

type VersionInfo struct {
	Version   string
	BuildDate time.Time
	Uptime    time.Duration
}

func VersionHandler(v *VersionInfo) http.HandlerFunc {
	t := time.Now()
	return func(w http.ResponseWriter, r *http.Request) {
		v.Uptime = time.Since(t)
		encoded, err := json.Marshal(v)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(encoded)
	}
}
