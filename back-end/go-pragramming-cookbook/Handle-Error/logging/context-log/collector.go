package context_log

import (
	"context"
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"os"
)

func Initialize() {
	log.SetHandler(text.New(os.Stdout))
	ctx := context.Background()

	ctx, e := FromContext(ctx, log.Log)

	e.Info("init")

	ctx = WithField(ctx, "test-id", "yj-choi")
	e.Info("starting")

	gatherName(ctx)
	e.Info("after gatherName")

	gatherLocation(ctx)
	e.Info("after gatherLocation")

}

func gatherName(ctx context.Context) {
	WithField(ctx, "name", "yoonjeong-choi-dev")
}

func gatherLocation(ctx context.Context) {
	WithFields(ctx, log.Fields{"city": "Gangnam", "state": "Seoul"})
}
