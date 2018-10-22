package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
)

// Graph structure
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

// NewGraph returns a new Graph.
func NewGraph() *Graph {
	return &Graph{
		NodeMap:                make(map[string]Node),
		EdgeMap:                make(map[string]Edge),
		IncomingNodeConnection: make(map[string]map[string]Node),
		OutgoingNodeConnection: make(map[string]map[string]Node),
	}
}

// ExistNode verify if node exist by a given id
func (g *Graph) ExistNode(id string) bool {
	_, ok := g.NodeMap[id]
	return ok
}

// GetNodeCount Return the current count of nodes in the Graph
func (g *Graph) GetNodeCount() int {
	return len(g.NodeMap)
}

// GetNode - Get a node in the Graph by id
func (g *Graph) GetNode(id string) Node {
	return g.NodeMap[id]
}

// AddNode - Add new node in the Graph
func (g Graph) AddNode(node Node) bool {

	if g.ExistNode(node.name) {
		return false
	}

	id := node.name
	g.NodeMap[id] = node
	return true
}

// DeleteNode - Delete a node by id
func (g *Graph) DeleteNode(id string) *Node {

	if !g.ExistNode(id) {
		return nil
	}

	node := g.NodeMap[id]

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

	return &node
}

// AddEdge - Add new Edge in the Graph
func (g *Graph) AddEdge(edge Edge) error {

	if !g.ExistNode(edge.begin.name) {
		return fmt.Errorf("%s does not exist in the graph", edge.begin)
	}
	if !g.ExistNode(edge.end.name) {
		return fmt.Errorf("%s does not exist in the graph", edge.end)
	}

	g.EdgeMap[edge.name] = edge

	if _, ok := g.OutgoingNodeConnection[edge.begin.name]; ok {
		g.OutgoingNodeConnection[edge.begin.name][edge.end.name] = edge.end
	} else {
		tmap := make(map[string]Node)
		tmap[edge.end.name] = edge.end
		g.OutgoingNodeConnection[edge.begin.name] = tmap
	}

	if _, ok := g.IncomingNodeConnection[edge.end.name]; ok {
		g.IncomingNodeConnection[edge.end.name][edge.begin.name] = edge.begin
	} else {
		tmap := make(map[string]Node)
		tmap[edge.begin.name] = edge.begin
		g.IncomingNodeConnection[edge.end.name] = tmap
	}

	return nil
}

// DeleteEdge - Delete a edge
func (g *Graph) DeleteEdge(edge Edge) Edge {

	removedEdge := g.EdgeMap[edge.name]

	firstMap := g.OutgoingNodeConnection[edge.begin.name]
	delete(firstMap, edge.end.name)

	secondMap := g.IncomingNodeConnection[edge.end.name]
	delete(secondMap, edge.begin.name)

	delete(g.EdgeMap, edge.name)

	return removedEdge
}

// WriteToFile - Write the entire graph as dot format in input file
func (g *Graph) WriteToFile(fileName string) {

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, g.String())
}

// Return a graphviz format string of all graph
func (g *Graph) String() string {
	var buffer bytes.Buffer

	s := fmt.Sprintf("digraph %s {\n", "Teste")
	buffer.WriteString(s)

	for _, node := range g.NodeMap {
		s = fmt.Sprintf("\t%s;\n", node.name)
		buffer.WriteString(s)
	}

	for _, edge := range g.EdgeMap {
		s = fmt.Sprintf("\t%s -> %s [label=%s, color=blue dir=none];\n", edge.begin.name, edge.end.name, edge.name)
		buffer.WriteString(s)
	}

	s = fmt.Sprintf("}\n")
	buffer.WriteString(s)

	return buffer.String()
}

// Verify if two nodes are adjacents
func (g *Graph) isAdjacent(nodeA Node, nodeB Node) bool {

	for _, edge := range g.EdgeMap {

		if edge.begin.name == nodeA.name && edge.end.name == nodeB.name {
			return true
		} else if edge.begin.name == nodeB.name && edge.end.name == nodeA.name {
			return true
		}

	}

	return false
}

// Print incoming connections
func (g *Graph) printIncomingConnections() {
	for key, value := range g.IncomingNodeConnection {
		fmt.Println("Incomming Map: ", key, value)
	}
}

// Print outgoung connections
func (g *Graph) printOutgoingConnections() {
	for key, value := range g.OutgoingNodeConnection {
		fmt.Println("Outgoing Map: ", key, value)
	}
}

func (g *Graph) floyd() {
	numberVertices := len(g.NodeMap)

	/* Created 2D map */
	shortestPath := make(map[string]map[string]float64, numberVertices)
	for _, value := range g.NodeMap {
		shortestPath[value.name] = make(map[string]float64, numberVertices)
	}

	/* Fill 2D map with 0's and infinities */
	for _, node := range g.NodeMap {
		for _, secondNode := range g.NodeMap {
			if node.name == secondNode.name {
				shortestPath[node.name][secondNode.name] = 0
			} else {
				shortestPath[node.name][secondNode.name] = math.Inf(0)
			}
		}
	}

	//fill shortestPath with current edge weight(u,v)
	for _, edge := range g.EdgeMap {
		shortestPath[edge.begin.name][edge.end.name] = edge.weight
	}

	for k := range shortestPath {
		for i := range shortestPath {
			for j := range shortestPath {

				if shortestPath[i][j] > (shortestPath[i][k] + shortestPath[k][j]) {
					shortestPath[i][j] = shortestPath[i][k] + shortestPath[k][j]
				}

			}
		}
	}

	fmt.Println(shortestPath)
}
