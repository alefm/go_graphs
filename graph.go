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

// newGraph returns a new Graph.
func NewGraph() *Graph {
	return &Graph{
		NodeMap:                make(map[string]Node),
		EdgeMap:                make(map[string]Edge),
		IncomingNodeConnection: make(map[string]map[string]Node),
		OutgoingNodeConnection: make(map[string]map[string]Node),
	}
}

// Verify if node exist by a given id
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
func (g Graph) AddNode(node Node) bool {

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
		if edge.begin.name == id || edge.end.name == id {
			delete(g.EdgeMap, edge.name)
		}
	}

	return true
}

// Add new Edge in the Graph
func (g *Graph) AddEdge(edge Edge) error {

	if !g.ExistNode(edge.begin.name) {
		return fmt.Errorf("%s does not exist in the graph.", edge.begin)
	}
	if !g.ExistNode(edge.end.name) {
		return fmt.Errorf("%s does not exist in the graph.", edge.end)
	}

	g.EdgeMap[edge.name] = edge

	if _, ok := g.IncomingNodeConnection[edge.begin.name]; ok {
		g.IncomingNodeConnection[edge.begin.name][edge.end.name] = edge.end
	} else {
		tmap := make(map[string]Node)
		tmap[edge.end.name] = edge.end
		g.IncomingNodeConnection[edge.begin.name] = tmap
	}

	if _, ok := g.OutgoingNodeConnection[edge.end.name]; ok {
		g.OutgoingNodeConnection[edge.end.name][edge.begin.name] = edge.begin
	} else {
		tmap := make(map[string]Node)
		tmap[edge.begin.name] = edge.begin
		g.OutgoingNodeConnection[edge.end.name] = tmap
	}
 	
	return nil
}

// Return a graphviz format string of all graph
func (g *Graph) String() string {

	fmt.Printf("digraph %s {\n", "Teste")

	for _, node := range g.NodeMap {
		fmt.Printf("\t%s;\n", node.name)
	}


	for _, edge := range g.EdgeMap {
		// fmt.Printf("Edge: %s begin %s end %s\n", edge.name, edge.begin.name, edge.end.name)	
		fmt.Printf("\t%s -> %s [label=%s, color=blue dir=none];\n", edge.begin.name, edge.end.name, edge.name)
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

// Verify if two nodes are adjacents
func (g *Graph) isAdjacent(nodeA Node, nodeB Node) bool {

	for _, edge := range g.EdgeMap {

		if edge.begin.name == nodeA.name && edge.end.name == nodeB.name{
			return true
		} else if edge.begin.name == nodeB.name && edge.end.name == nodeA.name {
			return true
		}

	}

	return false
}
