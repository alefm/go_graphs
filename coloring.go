package main

import (
	"sort"
)

// NodeDegree structure
type nodeDegree struct {
	node     Node
	colored  bool
	degree   int
	sliceIdx int
}

/* verify if a given Node have any connection with another node in list */
func (g *Graph) connectionInSlice(node Node, list []Node) bool {
	connection := false

	for _, nodeInList := range list {
		if nodeInList.GetName() == node.GetName() || g.isAdjacent(nodeInList, node) {
			connection = true
			break
		}
	}

	return connection
}

func isUncolored(list []nodeDegree) int {
	for i, value := range list {
		if !value.colored {
			return i
		}
	}

	return -1
}

// ColoringHeuristic should put colors in all nodes in graph
func (g *Graph) ColoringHeuristic() {
	pallete := [...]string{"gold", "green", "hotpink", "tan", "red", "blue", "tan", "yellow", "magenta", "cyan", "blueviolet", "olivedrab3"}
	colorIdx := 0
	neighbors := g.getNeighbors()
	colorMap := make(map[int][]Node)
	var degreeSize []nodeDegree

	for i, slice := range neighbors {
		degreeSize = append(degreeSize, nodeDegree{g.NodeList[i], false, len(slice), i})
	}

	// Reverse Sort by degree
	sort.Slice(degreeSize, func(i, j int) bool {
		return degreeSize[i].degree > degreeSize[j].degree
	})

	idx := isUncolored(degreeSize)
	for idx >= 0 {
		degreeSize[idx].node.SetColor(pallete[colorIdx])
		degreeSize[idx].colored = true
		g.NodeList[degreeSize[idx].sliceIdx] = degreeSize[idx].node

		// Set node in list of color
		var colorSlice []Node

		// Append node in color list
		colorSlice = append(colorSlice, degreeSize[idx].node)
		colorMap[colorIdx] = colorSlice

		for i := idx + 1; i < len(degreeSize); i++ {
			if !g.connectionInSlice(degreeSize[i].node, colorMap[colorIdx]) && !degreeSize[i].colored {
				degreeSize[i].node.SetColor(pallete[colorIdx])
				degreeSize[i].colored = true
				g.NodeList[degreeSize[i].sliceIdx] = degreeSize[i].node

				// Append node in color list
				colorSlice = append(colorSlice, degreeSize[i].node)
				colorMap[colorIdx] = colorSlice
			}
		}

		colorIdx = colorIdx + 1
		idx = isUncolored(degreeSize)
	}
}

// Coloring should put colors in all nodes in graph
func (g *Graph) Coloring() {
	pallete := [...]string{"gold", "green", "hotpink", "orchid", "red", "blue", "tan", "yellow", "magenta", "cyan", "blueviolet", "olivedrab3"}
	colorMap := make(map[int][]Node)
	colorIdx := 0

	var slice []Node
	colorMap[colorIdx] = slice

	for idx, nodeA := range g.NodeList {

		colored := false
		for i := 0; i <= colorIdx; i++ {
			if value, ok := colorMap[i]; ok {

				if !g.connectionInSlice(nodeA, value) {
					nodeA.SetColor(pallete[i])
					g.NodeList[idx] = nodeA
					value = append(value, nodeA)
					colorMap[i] = value
					colored = true
					break
				}

			}
		}

		if !colored {
			colorIdx = colorIdx + 1
			nodeA.SetColor(pallete[colorIdx])
			g.NodeList[idx] = nodeA
			var newSlice []Node
			newSlice = append(newSlice, nodeA)
			colorMap[colorIdx] = newSlice
		}
	}
}

func (g *Graph) ColoringFromPath() {
	g.ClearColors()
	for idx, node := range g.NodeList {
		for _, SearchPathNode := range g.SearchPath {
			if node.Name == SearchPathNode {
				node.SetColor("green")
				g.NodeList[idx] = node
			}
		}
	}
}

// Remove color to each nodes
func (g *Graph) ClearColors() {
	for idx, node := range g.NodeList {
		node.SetColor("")
		g.NodeList[idx] = node
	}
}
