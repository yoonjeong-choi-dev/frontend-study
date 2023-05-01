package main

import (
	"fmt"
	"github.com/trustmaster/goflow"
	goflow_example "goflow-example"
)

func main() {
	graph := goflow_example.NewEncodingApp()

	in := make(chan string)
	graph.SetInPort("In", in)

	// 그래프에서 데이터스트림을 처리 완료 여부 저장
	wait := goflow.Run(graph)
	for i := 0; i < 10; i++ {
		in <- fmt.Sprintf("Message-%d", i)
	}
	close(in)

	// 데이터스트림 처리까지 블로킹
	<-wait
}
