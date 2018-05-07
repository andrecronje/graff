package graff

// DirectedGraph is a graph supporting directed edges between nodes.
type DirectedGraph struct {
	*graph

	edges *directedEdgeList
}

// NewDirectedGraph creates a graph supporting directed edges between nodes.
func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{
		graph: newGraph(),
		edges: newDirectedEdgeList(),
	}
}

// Copy returns a new Directed Graph containing the same nodes and edges.
func (g *DirectedGraph) Copy() *DirectedGraph {
	return &DirectedGraph{
		graph: g.graph.Copy(),
		edges: g.edges.Copy(),
	}
}

func (g *DirectedGraph) EdgeCount() int {
	return g.edges.Count()
}

// AddEdge adds the edge to the graph.
func (g *DirectedGraph) AddEdge(from Node, to Node) {
	// prevent adding an edge referring to missing nodes
	if !g.NodeExists(from) {
		g.AddNode(from)
	}
	if !g.NodeExists(to) {
		g.AddNode(to)
	}

	g.edges.Add(from, to)
}

// RemoveEdge removes the edge from the graph.
func (g *DirectedGraph) RemoveEdge(from Node, to Node) {
	g.edges.Remove(from, to)
}

// EdgeExists checks whether the edge exists within the graph.
func (g *DirectedGraph) EdgeExists(from Node, to Node) bool {
	return g.edges.Exists(from, to)
}

func (g *DirectedGraph) HasPrecedingNodes(node Node) bool {
	return g.edges.HasPrecedingNodes(node)
}

func (g *DirectedGraph) PrecedingNodes(node Node) []Node {
	return g.edges.PrecedingNodes(node)
}

func (g *DirectedGraph) HasSucceedingNodes(node Node) bool {
	return g.edges.HasSucceedingNodes(node)
}

func (g *DirectedGraph) SucceedingNodes(node Node) []Node {
	return g.edges.SucceedingNodes(node)
}

func (g *DirectedGraph) RootNodes() []Node {
	rootNodes := make([]Node, 0)
	for _, node := range g.Nodes() {
		if !g.HasPrecedingNodes(node) {
			rootNodes = append(rootNodes, node)
		}
	}
	return rootNodes
}

type directedEdgeList struct {
	succeedingEdges map[Node]*nodeList
	precedingEdges  map[Node]*nodeList
}

func newDirectedEdgeList() *directedEdgeList {
	return &directedEdgeList{
		succeedingEdges: make(map[Node]*nodeList),
		precedingEdges:  make(map[Node]*nodeList),
	}
}

func (l *directedEdgeList) Copy() *directedEdgeList {
	succeedingEdges := make(map[Node]*nodeList, len(l.succeedingEdges))
	for node, edges := range l.succeedingEdges {
		succeedingEdges[node] = edges.Copy()
	}

	precedingEdges := make(map[Node]*nodeList, len(l.precedingEdges))
	for node, edges := range l.precedingEdges {
		precedingEdges[node] = edges.Copy()
	}

	return &directedEdgeList{
		succeedingEdges: succeedingEdges,
		precedingEdges:  precedingEdges,
	}
}

func (l *directedEdgeList) Count() int {
	return len(l.succeedingEdges)
}

func (l *directedEdgeList) HasSucceedingNodes(node Node) bool {
	_, ok := l.succeedingEdges[node]
	return ok
}

func (l *directedEdgeList) succeedingNodeList(node Node, create bool) *nodeList {
	if list, ok := l.succeedingEdges[node]; ok {
		return list
	}
	if create {
		list := newNodeList()
		l.succeedingEdges[node] = list
		return list
	}
	return nil
}

func (l *directedEdgeList) SucceedingNodes(node Node) []Node {
	if list := l.succeedingNodeList(node, false); list != nil {
		return list.Nodes()
	}
	return nil
}

func (l *directedEdgeList) HasPrecedingNodes(node Node) bool {
	_, ok := l.precedingEdges[node]
	return ok
}

func (l *directedEdgeList) precedingNodeList(node Node, create bool) *nodeList {
	if list, ok := l.precedingEdges[node]; ok {
		return list
	}
	if create {
		list := newNodeList()
		l.precedingEdges[node] = list
		return list
	}
	return nil
}

func (l *directedEdgeList) PrecedingNodes(node Node) []Node {
	if list := l.precedingNodeList(node, false); list != nil {
		return list.Nodes()
	}
	return nil
}

func (l *directedEdgeList) Add(from Node, to Node) {
	succeedingList := l.succeedingNodeList(from, true)
	succeedingList.Add(to)

	precedingList := l.precedingNodeList(to, true)
	precedingList.Add(from)
}

func (l *directedEdgeList) Remove(from Node, to Node) {
	if list := l.succeedingNodeList(from, false); list != nil {
		list.Remove(to)

		if list.Count() == 0 {
			delete(l.succeedingEdges, from)
		}
	}
	if list := l.precedingNodeList(to, false); list != nil {
		list.Remove(from)

		if list.Count() == 0 {
			delete(l.precedingEdges, to)
		}
	}
}

func (l *directedEdgeList) Exists(from Node, to Node) bool {
	if list := l.succeedingNodeList(from, false); list != nil {
		return list.Exists(to)
	}
	return false
}
