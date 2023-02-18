package main

import (
	"bytes"
	"dijkstra/graph"
	"fmt"
	"strconv"
	"text/tabwriter"
)

func DijkstraPathString(dist map[string]uint, prev map[string]string) string {
	buf := &bytes.Buffer{}
	writer := tabwriter.NewWriter(buf, 1, 5, 2, ' ', 0)
	writer.Write([]byte("Node\tDistance\tPrevious Node\t\n"))

	for node, value := range dist {
		writer.Write([]byte(node + "\t"))
		writer.Write([]byte(strconv.FormatUint(uint64(value), 10) + "\t"))
		writer.Write([]byte(prev[node] + "\t\n"))
	}

	writer.Flush()
	return buf.String()
}

func Example1() {
	g := graph.NewGraph()
	g.AddNodes("a", "b", "c", "d", "e")
	g.AddLink("a", "b", 6)

	g.AddLink("b", "d", 1)
	g.AddLink("b", "e", 2)

	g.AddLink("c", "b", 5)
	g.AddLink("c", "e", 5)

	g.AddLink("d", "a", 1)

	g.AddLink("e", "c", 4)
	g.AddLink("e", "d", 1)

	source := "a"
	dist, prev := g.Dijkstra(source)
	fmt.Println(DijkstraPathString(dist, prev))
}

func main() {
	Example1()
}
