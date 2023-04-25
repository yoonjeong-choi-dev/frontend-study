package pipeline_with_worker_pool

import "context"

// NewPipeline 채널들에 대한 파이프라인 생성
// : GoFlow 패키지의 그래프 생성하는 역할
// Reactive-Programming/goflow-example 참고
func NewPipeline(ctx context.Context, numEncoders, numPrinters int) (chan string, chan string) {
	inEncode := make(chan string, numEncoders)
	inPrint := make(chan string, numPrinters)
	outPrint := make(chan string, numPrinters)

	for i := 0; i < numEncoders; i++ {
		// Encode -> Printer 연결
		w := Worker{
			id:  i,
			in:  inEncode,
			out: inPrint,
		}
		go w.Work(ctx, Encode)
	}

	for i := 0; i < numPrinters; i++ {
		// Printer -> Out(클라이언트에서 사용)
		w := Worker{
			id:  i,
			in:  inPrint,
			out: outPrint,
		}

		go w.Work(ctx, Print)
	}

	// 데이터 스트림의 인풋 및 아웃풋 반환
	// => 해당 채널을 클라이언트에서 사용
	return inEncode, outPrint
}
