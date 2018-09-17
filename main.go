package main

import (
	"fmt"
)

func main() {
	graph := NewGraph()

	nodeA := Node{"A"}
	nodeB := Node{"B"}
	nodeC := Node{"C"}
	nodeD := Node{"D"}

	graph.AddNode(nodeA)
	graph.AddNode(nodeB)
	graph.AddNode(nodeC)
	graph.AddNode(nodeD)


	edge1 := Edge{"a", nodeA, nodeB, 15}
	edge2 := Edge{"b", nodeC, nodeD, 30}

	edge3 := Edge{"c", nodeB, nodeC, 15}
	edge4 := Edge{"d", nodeD, nodeA, 30}

	err := graph.AddEdge(edge1)
	if err != nil {
		fmt.Println(err)
	}

	err2 := graph.AddEdge(edge2)
	if err2 != nil {
		fmt.Println(err2)
	}

	err3 := graph.AddEdge(edge3)
	if err3 != nil {
		fmt.Println(err3)
	}

	err4 := graph.AddEdge(edge4)
	if err4 != nil {
		fmt.Println(err4)
	}

	graph.String()
}
