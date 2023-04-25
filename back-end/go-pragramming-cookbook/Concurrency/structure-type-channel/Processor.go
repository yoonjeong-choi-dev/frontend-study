package structure_type_channel

import "context"

func Processor(ctx context.Context, in chan *OpsRequest, out chan *OpsResponse) {
	for {
		select {
		case <-ctx.Done():
			return
		case r := <-in:
			out <- Process(r)
		}
	}
}
