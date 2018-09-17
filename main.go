package main

import (
	"fmt"
)

func (g *Graph) String() string {

	fmt.Printf("digraph %s {\n", "Teste")

	for _, node := range g.NodeMap {
		fmt.Printf("\t%s;\n", node.name)
	}


	for _, edge := range g.EdgeMap {
		// fmt.Printf("Edge: %s src %s dst %s\n", edge.name, edge.src.name, edge.dst.name)	
		fmt.Printf("\t%s -> %s [label=%s, color=red];\n", edge.src.name, edge.dst.name, edge.name)
	}

	/*for key, sublist := range g.IncomingNodeConnection {
		fmt.Printf("IncKey %s: ", key)

		for _, node := range sublist {
			fmt.Printf("%s ", node.name)
		}

		fmt.Printf("\n")
	}

	for key, sublist := range g.OutgoingNodeConnection {
		fmt.Printf("OutKey %s: ", key)

		for _, node := range sublist {
			fmt.Printf("%s  ", node.name)
		}

		fmt.Printf("\n")
	}*/

	fmt.Printf("}\n")

	return "l"
}

func main() {
	graph := NewGraph()

	graph.AddNode(*NewNode("A"))
	graph.AddNode(*NewNode("B"))
	graph.AddNode(*NewNode("C"))
	graph.AddNode(*NewNode("D"))

	nodeA := graph.GetNode("A")
	nodeB := graph.GetNode("B")
	nodeC := graph.GetNode("C")
	nodeD := graph.GetNode("D")

	edge := NewEdge("E", nodeA, nodeB, 64)
	edge2 := NewEdge("F", nodeC, nodeD, 64)

	err := graph.AddEdge(*edge)
	if err != nil {
		fmt.Println(err)
	}

	err2 := graph.AddEdge(*edge2)
	if err2 != nil {
		fmt.Println(err2)
	}

	graph.String()
}
