package main

import (
	"math"
)

// Floyd algorithm
func (g *Graph) Floyd() ([][]float64, [][]string) {
	numberVertices := len(g.NodeList)

	/* Created 2D slice */
	shortestPath := make([][]float64, numberVertices)
	predecessor := make([][]string, numberVertices)

	for index := range shortestPath {
		shortestPath[index] = make([]float64, numberVertices)
		predecessor[index] = make([]string, numberVertices)
	}

	/* Fill 2D slice predecessor */
	for _, value := range g.EdgeList {
		predecessor[g.ExistNode(value.begin.Name)][g.ExistNode(value.end.Name)] = value.end.Name
	}

	/* Fill 2D slice */
	for i := 0; i < len(g.NodeList); i++ {
		for j := 0; j < len(g.NodeList); j++ {
			if index := g.GetEdgeIndex(g.NodeList[i].Name, g.NodeList[j].Name); index >= 0 {
				shortestPath[i][j] = g.EdgeList[index].weight
				predecessor[i][j] = g.NodeList[j].GetName()
			} else if i == j {
				shortestPath[i][j] = 0
			} else {
				shortestPath[i][j] = math.Inf(0)
			}
		}
	}

	for k := 0; k < len(shortestPath); k++ {
		for i := 0; i < len(shortestPath); i++ {
			for j := 0; j < len(shortestPath); j++ {
				if shortestPath[i][k] < math.Inf(0) && shortestPath[k][j] < math.Inf(0) {
					if shortestPath[i][j] > (shortestPath[i][k] + shortestPath[k][j]) {
						shortestPath[i][j] = shortestPath[i][k] + shortestPath[k][j]
						predecessor[i][j] = predecessor[i][k]
					}
				}
			}
		}
	}

	return shortestPath, predecessor
}

// FloydPath find the best path to floyd algorithm
func (g *Graph) FloydPath(predecessor [][]string, begin string, end string) []string {
	path := make([]string, 0)

	path = append(path, begin)
	for begin != end {
		beginIdx := g.ExistNode(begin)
		endIdx := g.ExistNode(end)
		begin = predecessor[beginIdx][endIdx]
		path = append(path, begin)
	}

	return path
}
