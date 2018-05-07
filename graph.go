package graff

type graph struct {
	nodes *nodeList
}

func newGraph() *graph {
	return &graph{
		nodes: newNodeList(),
	}
}

func (g *graph) Copy() *graph {
	return &graph{
		nodes: g.nodes.Copy(),
	}
}

func (g *graph) Nodes() []Node {
	return g.nodes.Nodes()
}

func (g *graph) NodeCount() int {
	return g.nodes.Count()
}

func (g *graph) AddNode(node Node) {
	g.AddNodes(node)
}

func (g *graph) AddNodes(nodes ...Node) {
	g.nodes.Add(nodes...)
}

func (g *graph) RemoveNode(node Node) {
	g.RemoveNodes(node)
}

func (g *graph) RemoveNodes(nodes ...Node) {
	g.nodes.Remove(nodes...)
}

func (g *graph) NodeExists(node Node) bool {
	return g.nodes.Exists(node)
}
