package main

import (
	"fmt"
)

func main() {
	graph := NewGraph()

	graph.AddNode(Node{"A"})
	graph.AddNode(Node{"B"})
	graph.AddNode(Node{"C"})
	graph.AddNode(Node{"D"})

	nodeA := graph.GetNode("A")
	nodeB := graph.GetNode("B")
	nodeC := graph.GetNode("C")
	nodeD := graph.GetNode("D")

	edge := Edge{"E", nodeA, nodeB, 64}
	edge2 := Edge{"F", nodeC, nodeD, 64}

	err := graph.AddEdge(edge)
	if err != nil {
		fmt.Println(err)
	}

	err2 := graph.AddEdge(edge2)
	if err2 != nil {
		fmt.Println(err2)
	}

	graph.String()


}
