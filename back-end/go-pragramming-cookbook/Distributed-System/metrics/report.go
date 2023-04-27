package metrics

import (
	gm "github.com/rcrowley/go-metrics"
	"net/http"
)

func ReportHandler(w http.ResponseWriter, r *http.Request) {
	t := gm.GetOrRegisterTimer("reporthandler.write", nil)

	w.WriteHeader(http.StatusOK)
	t.Time(func() {
		gm.WriteJSONOnce(gm.DefaultRegistry, w)
	})
}
