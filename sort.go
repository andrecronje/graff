package graff

import (
	"container/list"
	"errors"
)

// Errors relating to the DFSSorter.
var (
	ErrCyclicGraph = errors.New("The graph cannot be cyclic")
)

// DFSSorter topologically sorts a directed graph's nodes based on the
// directed edges between them using the Depth-first search algorithm.
type DFSSorter struct {
	graph                *DirectedGraph
	nodes                []Node
	undiscovered         *list.List
	undiscoveredElements map[Node]*list.Element
	visiting             map[Node]bool
	discovered           map[Node]bool
}

// NewDFSSorter creates a topological sorter for the specified directed
// graph's nodes based on the directed edges between them using the
// Depth-first search algorithm.
func NewDFSSorter(graph *DirectedGraph) *DFSSorter {
	return &DFSSorter{
		graph: graph,
	}
}

func (s *DFSSorter) init() {
	s.nodes = make([]Node, 0, s.graph.NodeCount())
	s.visiting = make(map[Node]bool)
	s.discovered = make(map[Node]bool, s.graph.NodeCount())

	s.undiscovered = list.New()
	s.undiscoveredElements = make(map[Node]*list.Element, s.graph.NodeCount())
	for _, node := range s.graph.Nodes() {
		element := s.undiscovered.PushFront(node)
		s.undiscoveredElements[node] = element
	}

}

// Sort returns the nodes in topological order.
func (s *DFSSorter) Sort() ([]Node, error) {
	s.init()

	// > while there are unmarked nodes do
	for s.undiscovered.Len() > 0 {
		for e := s.undiscovered.Front(); e != nil; e = e.Next() {
			node := e.Value.(Node)
			if err := s.visit(node); err != nil {
				return nil, err
			}
		}
	}

	// as the nodes were appended to the slice for performance reasons,
	// rather than prepended as correctly stated by the algorithm,
	// we need to reverse the sorted slice
	for i, j := 0, len(s.nodes)-1; i < j; i, j = i+1, j-1 {
		s.nodes[i], s.nodes[j] = s.nodes[j], s.nodes[i]
	}

	return s.nodes, nil
}

func (s *DFSSorter) visit(node Node) error {
	// > if n has a permanent mark then return
	if discovered, ok := s.discovered[node]; ok && discovered {
		return nil
	}
	// > if n has a temporary mark then stop (not a DAG)
	if visiting, ok := s.visiting[node]; ok && visiting {
		return ErrCyclicGraph
	}

	// > mark n temporarily
	s.visiting[node] = true

	// for each node m with an edge from n to m do
	succeeding := s.graph.SucceedingNodes(node)
	for i := len(succeeding) - 1; i >= 0; i-- {
		if err := s.visit(succeeding[i]); err != nil {
			return err
		}
	}

	s.discovered[node] = true
	delete(s.visiting, node)

	if element, ok := s.undiscoveredElements[node]; ok {
		s.undiscovered.Remove(element)
		delete(s.undiscoveredElements, node)
	}

	s.nodes = append(s.nodes, node)
	return nil
}

// DFSSort returns the graph's nodes in topological order based on the
// directed edges between them using the Depth-first search algorithm.
func (g *DirectedGraph) DFSSort() ([]Node, error) {
	sorter := NewDFSSorter(g)
	return sorter.Sort()
}
