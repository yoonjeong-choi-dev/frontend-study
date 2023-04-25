package pipeline_with_worker_pool

import "context"

// Job 워커가 수행할 작업
type Job string

const (
	Print  Job = "print"
	Encode Job = "encode"
)

// Worker 데이터 스트림의 input 및 output 채널
// => 데이터 스트림의 각 노드 역할
type Worker struct {
	id  int
	in  chan string
	out chan string
}

func (w *Worker) Work(ctx context.Context, job Job) {
	switch job {
	case Print:
		w.Print(ctx)
	case Encode:
		w.Encode(ctx)
	default:
		return
	}
}
