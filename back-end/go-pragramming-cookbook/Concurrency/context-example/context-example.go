package context_example

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func Example() {
	ctx := context.Background()
	ctx = Setup(ctx)

	rand.Seed(time.Now().UnixNano())

	timeoutCtx, timeoutCancel := context.WithTimeout(
		ctx,
		time.Duration(rand.Intn(2))*time.Millisecond,
	)
	defer timeoutCancel()

	deadlineCtx, deadlineCancel := context.WithDeadline(
		ctx,
		time.Now().Add(time.Duration(rand.Intn(2))*time.Millisecond),
	)
	defer deadlineCancel()

	for {
		select {
		case <-timeoutCtx.Done():
			fmt.Println(GetValueFromContext(ctx, timeoutKey))
			return
		case <-deadlineCtx.Done():
			fmt.Println(GetValueFromContext(ctx, deadlineKey))
			return
		}
	}
}
