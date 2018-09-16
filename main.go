package main

import (
	"fmt"
)

type Graph struct {

	//Store all nodes in the graph
	NodeMap map[string]Node

	//Store all edges in the graph
	EdgeMap map[string]Edge

	// IncomingNodeConnection maps its Node to incoming Nodes with its edge weight (incoming edges to its Node).
	IncomingNodeConnection map[string]map[string]Node

	// OutgoingNodeConnection maps its Node to outgoing Nodes with its edge weight (outgoing edges from its Node).
	OutgoingNodeConnection map[string]map[string]Node
}

type Edge struct {
	// Unique ID
	name   string
	src    Node
	dst    Node
	weight float64
}

type Node struct {
	// Unique ID
	name string
}

// newGraph returns a new Graph.
func NewGraph() *Graph {
	return &Graph{
		NodeMap:                make(map[string]Node),
		EdgeMap:                make(map[string]Edge),
		IncomingNodeConnection: make(map[string]map[string]Node),
		OutgoingNodeConnection: make(map[string]map[string]Node),
	}
}

// NewEdge returns a new Edge.
func NewEdge(id string, src, dst Node, weight float64) *Edge {
	return &Edge{
		name:   id,
		src:    src,
		dst:    dst,
		weight: weight,
	}
}

// NewNode returns a new Node.
func NewNode(id string) *Node {
	return &Node{
		name: id,
	}
}

func (g *Graph) ExistNode(id string) bool {
	_, ok := g.NodeMap[id]
	return ok
}

// Return the current count of nodes in the Graph
func (g *Graph) GetNodeCount() int {
	return len(g.NodeMap)
}

// Get a node in the Graph by id
func (g *Graph) GetNode(id string) Node {
	return g.NodeMap[id]
}

// Add new node in the Graph
func (g *Graph) AddNode(node Node) bool {

	if g.ExistNode(node.name) {
		return false
	}

	id := node.name
	g.NodeMap[id] = node
	return true
}

// Delete a node by id
func (g *Graph) DeleteNode(id string) bool {

	if !g.ExistNode(id) {
		return false
	}

	delete(g.NodeMap, id)

	delete(g.IncomingNodeConnection, id)
	for _, submap := range g.IncomingNodeConnection {
		delete(submap, id)
	}

	delete(g.OutgoingNodeConnection, id)
	for _, submap := range g.OutgoingNodeConnection {
		delete(submap, id)
	}

	for _, edge := range g.EdgeMap {
		if edge.src.name == id || edge.dst.name == id {
			delete(g.EdgeMap, edge.name)
		}
	}

	return true
}

// Add new Edge in the Graph
func (g *Graph) AddEdge(edge Edge) error {

	if !g.ExistNode(edge.src.name) {
		return fmt.Errorf("%s does not exist in the graph.", edge.src)
	}
	if !g.ExistNode(edge.dst.name) {
		return fmt.Errorf("%s does not exist in the graph.", edge.dst)
	}

	id := edge.name
	g.EdgeMap[id] = edge

	if _, ok := g.IncomingNodeConnection[edge.src.name]; ok {
		g.IncomingNodeConnection[edge.src.name][edge.dst.name] = edge.dst
	} else {
		tmap := make(map[string]Node)
		tmap[edge.dst.name] = edge.dst
		g.IncomingNodeConnection[edge.src.name] = tmap
	}

	if _, ok := g.OutgoingNodeConnection[edge.dst.name]; ok {
		g.OutgoingNodeConnection[edge.dst.name][edge.src.name] = edge.src
	} else {
		tmap := make(map[string]Node)
		tmap[edge.src.name] = edge.src
		g.OutgoingNodeConnection[edge.dst.name] = tmap
	}
 	
	return nil
}

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
