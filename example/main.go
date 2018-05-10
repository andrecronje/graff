package main

import (
	"fmt"
	"log"

	"github.com/jackwakefield/graff"
)

func main() {
	//   		 +---+
	//           | A |
	//           +---+
	//             |
	//             |
	//             v
	// +---+     +---+
	// | B | --> | D |
	// +---+     +---+
	//   |         |
	//   |         |
	//   v         v
	// +---+     +-------------+     +---+
	// | C | --> |      E      | --> | F |
	// +---+     +-------------+     +---+
	//             |    |              |
	//             |    |              |
	//             v    |              |
	//           +---+  |              |
	//        +- | H |  |              |
	//        |  +---+  |              |
	//        |    |    |              |
	//        |    |    |              |
	//        |    v    |              |
	//        |  +---+  |              |
	//        |  | I | <+              |
	//        |  +---+                 |
	//        |    |                   |
	//        |    |                   |
	//        |    v                   |
	//        |  +-------------+       |
	//        |  |      G      | <-----+
	//        |  +-------------+
	//        |    ^         ^
	//        |    |         |
	//        |    |         |
	//        |  +---+       |
	//        |  | J |       |
	//        |  +---+       |
	//        |    |         |
	//        |    |         |
	//        |    v         |
	//        |  +---+       |
	//        +> | K | ------+
	//           +---+

	graph := graff.NewDirectedGraph()
	graph.AddNodes("A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K")

	graph.AddEdge("A", "D")
	graph.AddEdge("B", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "E")
	graph.AddEdge("D", "E")
	graph.AddEdge("E", "F")
	graph.AddEdge("E", "I")
	graph.AddEdge("E", "H")
	graph.AddEdge("F", "G")
	graph.AddEdge("H", "I")
	graph.AddEdge("H", "K")
	graph.AddEdge("I", "G")
	graph.AddEdge("J", "G")
	graph.AddEdge("J", "K")
	graph.AddEdge("K", "G")

	dfs, err := graph.DFSSort()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Depth-first search sort:", dfs)

	coffmanGraham, err := graph.CoffmanGrahamSort(4)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Coffman-Graham sort (W=4):", coffmanGraham)
}
