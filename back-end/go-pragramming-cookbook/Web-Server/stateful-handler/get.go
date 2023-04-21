package stateful_handler

import (
	"encoding/json"
	"net/http"
)

// GetValue : Wrapping http.HandleFunc by closure
// 핸들러에 상태를 추가하는 방법 1: 클로저를 통해 상태 전달
// => 핸들러 자체를 반환하는 방식
func (c *Controller) GetValue(useDefault bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		value := "default"
		if !useDefault {
			value = c.storage.Get()
		}

		w.WriteHeader(http.StatusOK)
		p := Payload{Value: value}
		if payload, err := json.Marshal(p); err == nil {
			w.Write(payload)
		}
	}
}
