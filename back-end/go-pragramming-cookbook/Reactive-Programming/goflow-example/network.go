package goflow_example

import "github.com/trustmaster/goflow"

// NewEncodingApp 정의한 컴포넌트들을 연결하는 네트워크(그래프) 생성
func NewEncodingApp() *goflow.Graph {
	graph := goflow.NewGraph()

	// 컴포넌트 타입 정의
	// => define vertices
	graph.Add("encoder", new(Encoder))
	graph.Add("printer", new(Printer))

	// 컴포넌트 연결
	// => define edges via channel
	graph.Connect("encoder", "ResultMessage", "printer", "Line")

	// 컴포넌트 외부 진입점(entry point) 정의
	// "In" 에 대한 설정은 외부 클라이언트 코드에서 정의
	graph.MapInPort("In", "encoder", "Val")

	return graph
}
