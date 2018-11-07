package main

import "math"

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

	/* Fill 2D slice */
	for i := 0; i < len(g.NodeList); i++ {
		for j := 0; j < len(g.NodeList); j++ {
			if index := g.ExistEdge(g.NodeList[i].Name, g.NodeList[j].Name); index >= 0 {
				shortestPath[i][j] = g.EdgeList[index].weight
			} else if i == j {
				shortestPath[i][j] = 0
			} else {
				shortestPath[i][j] = math.Inf(0)
			}
		}
	}

	/* Fill 2D slice predecessor */
	for _, value := range g.EdgeList {
		predecessor[g.ExistNode(value.begin.Name)][g.ExistNode(value.end.Name)] = value.end.Name
	}

	for k := 0; k < len(shortestPath); k++ {
		for i := 0; i < len(shortestPath); i++ {
			for j := 0; j < len(shortestPath); j++ {

				if shortestPath[i][j] > (shortestPath[i][k] + shortestPath[k][j]) {
					shortestPath[i][j] = shortestPath[i][k] + shortestPath[k][j]
					predecessor[i][j] = predecessor[i][k]
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
		begin = predecessor[g.ExistNode(begin)][g.ExistNode(end)]
		path = append(path, begin)
	}

	return path
}
