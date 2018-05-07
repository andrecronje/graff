package graff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTestDirectedGraph() *DirectedGraph {
	graph := NewDirectedGraph()
	graph.AddNodes("A", "B", "C", "D")
	return graph
}

func TestNewDirectedGraph(t *testing.T) {
	graph := NewDirectedGraph()
	assert.NotNil(t, graph, "graph should not be nil")
	assert.Zero(t, graph.NodeCount(), "graph.NodeCount() should equal zero")
	assert.Empty(t, graph.Nodes(), "graph.Nodes() should equal empty")
	assert.Zero(t, graph.EdgeCount(), "graph.EdgeCount() should equal zero")
}

func TestDirectedGraphAddEdge(t *testing.T) {
	graph := newTestDirectedGraph()
	graph.AddEdge("A", "B")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "B")

	assert.Equal(t, 3, graph.EdgeCount(), "graph.EdgeCount() should equal 3")
	assert.True(t, graph.EdgeExists("A", "B"), "graph.EdgeExists(A, B) should equal true")
	assert.True(t, graph.EdgeExists("B", "D"), "graph.EdgeExists(B, D) should equal true")
	assert.True(t, graph.EdgeExists("C", "B"), "graph.EdgeExists(C, B) should equal true")
}

func TestDirectedGraphAddEdgeDuplicate(t *testing.T) {
	graph := newTestDirectedGraph()
	graph.AddEdge("A", "B")
	graph.AddEdge("B", "C")
	graph.AddEdge("B", "C")

	assert.Equal(t, 2, graph.EdgeCount(), "graph.EdgeCount() should equal 2")
	assert.True(t, graph.EdgeExists("A", "B"), "graph.EdgeExists(A, B) should equal true")
	assert.True(t, graph.EdgeExists("B", "C"), "graph.EdgeExists(B, C) should equal true")
}

func TestDirectedGraphAddEdgeMissingNodes(t *testing.T) {
	graph := NewDirectedGraph()
	graph.AddEdge("A", "B")
	graph.AddEdge("B", "C")

	assert.Equal(t, 3, graph.NodeCount(), "graph.NodeCount() should equal 2")
	assert.Equal(t, 2, graph.EdgeCount(), "graph.EdgeCount() should equal 2")
	assert.True(t, graph.EdgeExists("A", "B"), "graph.EdgeExists(A, B) should equal true")
	assert.True(t, graph.EdgeExists("B", "C"), "graph.EdgeExists(B, C) should equal true")
}

func TestDirectedGraphRemoveEdge(t *testing.T) {
	graph := newTestDirectedGraph()
	graph.AddEdge("A", "B")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "B")
	graph.RemoveEdge("A", "B")
	graph.RemoveEdge("C", "B")

	assert.Equal(t, 1, graph.EdgeCount(), "graph.EdgeCount() should equal 1")
	assert.False(t, graph.EdgeExists("A", "B"), "graph.EdgeExists(A, B) should equal false")
	assert.False(t, graph.EdgeExists("C", "B"), "graph.EdgeExists(C, B) should equal false")
	assert.True(t, graph.EdgeExists("B", "D"), "graph.EdgeExists(B, D) should equal true")
}

func TestDirectedGraphRemoveEdgeMissing(t *testing.T) {
	graph := newTestDirectedGraph()
	graph.AddEdge("C", "B")
	graph.RemoveEdge("D", "A")
	graph.RemoveEdge("C", "B")

	assert.Zero(t, graph.EdgeCount(), "graph.EdgeCount() should equal zero")
}

func TestDirectedGraphRootNodes(t *testing.T) {
	graph := newTestDirectedGraph()
	graph.AddEdge("A", "B")
	graph.AddEdge("B", "C")
	graph.AddEdge("D", "C")
	graph.AddEdge("E", "C")
	graph.AddEdge("F", "E")

	assert.Equal(t, []Node{"A", "D", "F"}, graph.RootNodes(), "graph.RootNodes() should equal [A, D, F]")
}
