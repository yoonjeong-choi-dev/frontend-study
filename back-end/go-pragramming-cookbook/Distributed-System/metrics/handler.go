package metrics

import (
	gm "github.com/rcrowley/go-metrics"
	"net/http"
	"time"
)

func CounterHandler(w http.ResponseWriter, r *http.Request) {
	c := gm.GetOrRegisterCounter("counterhandler.counter", nil)
	c.Inc(1)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success to increase counter"))
}

func TimerHandler(w http.ResponseWriter, r *http.Request) {
	cur := time.Now()
	t := gm.GetOrRegisterTimer("timerhandler.timer", nil)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success to update time"))

	t.UpdateSince(cur)
}
