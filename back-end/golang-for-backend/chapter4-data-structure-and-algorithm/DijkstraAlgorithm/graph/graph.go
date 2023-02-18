package graph

const IntInfinity = ^uint(0)

type NodeName = string

type Node struct {
	Name  NodeName
	links []Edge
}

type Edge struct {
	from   *Node
	to     *Node
	weight uint
}

type DirectedWeightedGraph struct {
	nodes map[NodeName]*Node
}

func NewGraph() *DirectedWeightedGraph {
	return &DirectedWeightedGraph{nodes: map[NodeName]*Node{}}
}

func (graph *DirectedWeightedGraph) AddNodes(names ...NodeName) {
	for _, name := range names {
		if _, ok := graph.nodes[name]; !ok {
			graph.nodes[name] = &Node{Name: name, links: []Edge{}}
		}
	}
}

func (graph *DirectedWeightedGraph) AddLink(from, to NodeName, weight int) {
	fromNode := graph.nodes[from]
	toNode := graph.nodes[to]

	// Undirected
	uWeight := uint(weight)
	fromNode.links = append(fromNode.links, Edge{from: fromNode, to: toNode, weight: uWeight})
}

type Graph = DirectedWeightedGraph

func (graph *Graph) Dijkstra(source string) (map[string]uint, map[string]string) {
	// 각 노드에 대한 경로 비용 및 경로에서의 이전 노드
	dist, prev := map[string]uint{}, map[string]string{}

	for _, node := range graph.nodes {
		dist[node.Name] = IntInfinity
		prev[node.Name] = ""
	}

	visited := map[string]bool{}
	dist[source] = 0

	// source 에서 시작하여 가장 짧은 path를 가지는 노드부터 탐색
	// getClosestNonVisitedNode 에서 "" 노드 반환 => 모든 노드를 방문했거나 비용 계산이 완료되었다는 의미
	for cur := source; cur != ""; cur = getClosestNonVisitedNode(dist, visited) {
		curDist := dist[cur]

		// 현재 노드와 연결된 노드들 탐색
		for _, link := range graph.nodes[cur].links {
			// 이미 방문한 노드면 무시
			if _, ok := visited[link.to.Name]; ok {
				continue
			}

			nextDist := curDist + link.weight
			next := link.to.Name
			if nextDist < dist[next] {
				dist[next] = nextDist
				prev[next] = cur
			}
		}
		visited[cur] = true
	}

	return dist, prev
}

func getClosestNonVisitedNode(dist map[string]uint, visited map[string]bool) string {
	lowestCost := IntInfinity
	lowestNode := ""

	for key, d := range dist {
		// 이미 방문한 경우나 아직 경로가 계산되지 않은 경우 무시
		if _, ok := visited[key]; ok || d == IntInfinity {
			continue
		}

		// 기존에 계산한 경로의 비용보다 작은 경우 업데이트
		if d < lowestCost {
			lowestCost = d
			lowestNode = key
		}
	}

	return lowestNode
}
