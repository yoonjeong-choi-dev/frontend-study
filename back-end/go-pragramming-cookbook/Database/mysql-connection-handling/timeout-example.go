package connectionpools

import (
	"context"
	"time"
)

func ExecWithTimeout() error {
	db, err := Setup()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now())
	defer cancel()

	_, err = db.BeginTx(ctx, nil)
	return err
}
