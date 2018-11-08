package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

// Graph structure
type Graph struct {

	//Store all nodes in the graph
	NodeList []Node

	//Store all edges in the graph
	EdgeList []Edge
}

// NewGraph returns a new Graph.
func NewGraph() *Graph {
	return &Graph{
		NodeList: make([]Node, 0),
		EdgeList: make([]Edge, 0),
	}
}

// ExistNode verify if node exist by a given id and return current slice index
func (g *Graph) ExistNode(id string) int {
	for key, value := range g.NodeList {
		if value.Name == id {
			return key
		}
	}
	return -1
}

// ExistEdge verify if edge exist by a given begin and end id
func (g *Graph) ExistEdge(begin string, end string) int {
	for key, value := range g.EdgeList {
		if value.begin.Name == begin && value.end.Name == end {
			return key
		}
	}
	return -1
}

// GetEdgeIndex return edge index
func (g *Graph) GetEdgeIndex(begin string, end string) int {
	for key, value := range g.EdgeList {
		if value.begin.Name == begin && value.end.Name == end ||
			value.begin.Name == end && value.end.Name == begin {
			return key
		}
	}
	return -1
}

// GetNodeCount Return the current count of nodes in the Graph
func (g *Graph) GetNodeCount() int {
	return len(g.NodeList)
}

// GetNode - Get a node in the Graph by id
func (g *Graph) GetNode(id string) *Node {
	for _, value := range g.NodeList {
		if value.Name == id {
			return &value
		}
	}
	return nil
}

// GetEdge - Get a edge in the Graph by id
func (g *Graph) GetEdge(id string) *Edge {
	for _, value := range g.EdgeList {
		if value.name == id {
			return &value
		}
	}
	return nil
}

// AddNode - Add new node in the Graph
func (g *Graph) AddNode(node Node) bool {

	if g.ExistNode(node.Name) >= 0 {
		return false
	}

	g.NodeList = append(g.NodeList, node)
	return true
}

// DeleteNode - Delete a node by id
func (g *Graph) DeleteNode(id string) *Node {

	if g.ExistNode(id) < 0 {
		return nil
	}

	node := g.GetNode(id)

	// Removing node from slice
	index := g.ExistNode(id)
	g.NodeList = append(g.NodeList[:index], g.NodeList[index+1:]...)

	for key, edge := range g.EdgeList {
		if edge.begin.Name == id || edge.end.Name == id {
			// Removing edge from slice
			g.EdgeList = append(g.EdgeList[:key], g.EdgeList[key+1:]...)
		}
	}

	return node
}

// AddEdge - Add new Edge in the Graph
func (g *Graph) AddEdge(edge Edge) error {
	if g.ExistNode(edge.begin.Name) < 0 {
		return fmt.Errorf("%s does not exist in the graph", edge.begin.Name)
	}
	if g.ExistNode(edge.end.Name) < 0 {
		return fmt.Errorf("%s does not exist in the graph", edge.end.Name)
	}

	g.EdgeList = append(g.EdgeList, edge)

	return nil
}

// DeleteEdge - Delete a edge
func (g *Graph) DeleteEdge(edge Edge) *Edge {

	index := g.ExistEdge(edge.begin.Name, edge.end.Name)

	if index < 0 {
		return nil
	}

	removedEdge := g.EdgeList[index]

	g.EdgeList = append(g.EdgeList[:index], g.EdgeList[index+1:]...)

	return &removedEdge
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

	for _, node := range g.NodeList {
		if node.Color != "" {
			s = fmt.Sprintf("\t%s [fillcolor=%s,style=filled];\n", node.Name, node.GetColor())
		} else {
			s = fmt.Sprintf("\t%s [color=%s];\n", node.Name, node.GetColor())
		}

		buffer.WriteString(s)
	}

	for _, edge := range g.EdgeList {
		s = fmt.Sprintf("\t%s -> %s [label=%.2f, color=%s, dir=none];\n", edge.begin.Name, edge.end.Name, edge.weight, edge.GetColor())
		buffer.WriteString(s)
	}

	s = fmt.Sprintf("}\n")
	buffer.WriteString(s)

	return buffer.String()
}

// Verify if two nodes are adjacents
func (g *Graph) isAdjacent(nodeA Node, nodeB Node) bool {

	for _, edge := range g.EdgeList {

		if edge.begin.Name == nodeA.Name && edge.end.Name == nodeB.Name {
			return true
		} else if edge.begin.Name == nodeB.Name && edge.end.Name == nodeA.Name {
			return true
		}

	}

	return false
}
