package raft_concensus

import (
	"fmt"
	"net/http"
	"time"
)

func StateTransitionHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	state := r.FormValue("next")
	for address, raft := range rafts {
		// 리더가 아닌 경우 무시
		if address != raft.Leader() {
			continue
		}

		// 리더 raft 가 다음 상태로 변경 시도
		result := raft.Apply([]byte(state), 1*time.Second)
		if result.Error() != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newState, ok := result.Response().(string)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if newState != state {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid transition"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("-> %s", newState)))
	}
}
