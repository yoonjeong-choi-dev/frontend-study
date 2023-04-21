package stateful_handler

import (
	"encoding/json"
	"net/http"
)

// SetValue implement http.HandleFunc
// 핸들러에 상태를 추가하는 방법 2: 구조체의 메서드로 핸들러 구현
func (c *Controller) SetValue(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	value := r.FormValue("value")
	c.storage.Put(value)

	w.WriteHeader(http.StatusOK)
	p := Payload{Value: value}
	if payload, err := json.Marshal(p); err == nil {
		w.Write(payload)
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
