package graff

// Node represents a graph node.
type Node = interface{}

type nodeList struct {
	nodes []Node
}

func newNodeList() *nodeList {
	return &nodeList{
		nodes: make([]Node, 0),
	}
}

func (l *nodeList) Copy() *nodeList {
	nodes := make([]Node, len(l.nodes))
	copy(nodes, l.nodes)

	return &nodeList{
		nodes: nodes,
	}
}

func (l *nodeList) Nodes() []Node {
	return l.nodes
}

func (l *nodeList) Count() int {
	return len(l.nodes)
}

func (l *nodeList) Exists(node Node) bool {
	_, ok := l.indexOf(node)
	return ok
}

func (l *nodeList) indexOf(node Node) (index int, ok bool) {
	for index, item := range l.nodes {
		if item == node {
			return index, true
		}
	}
	return 0, false
}

func (l *nodeList) Add(nodes ...Node) {
	for _, node := range nodes {
		if l.Exists(node) {
			return
		}

		l.nodes = append(l.nodes, node)
	}
}

func (l *nodeList) Remove(nodes ...Node) {
	for _, node := range nodes {
		index, ok := l.indexOf(node)
		if !ok {
			return
		}

		copy(l.nodes[index:], l.nodes[index+1:])
		l.nodes[len(l.nodes)-1] = nil
		l.nodes = l.nodes[:len(l.nodes)-1]
	}
}
