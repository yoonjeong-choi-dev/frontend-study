// Package kafka_with_goflow goflow 그래프 정의
package kafka_with_goflow

import "github.com/trustmaster/goflow"

func NewUpperApp() *goflow.Graph {
	graph := goflow.NewGraph()

	// Add vertices
	graph.Add("upper", new(Upper))
	graph.Add("printer", new(Printer))

	// And Edges
	graph.Connect("upper", "Result", "printer", "Line")

	// And Entry Vertex
	graph.MapInPort("In", "upper", "Val")

	return graph
}
