package main

import "math"

// Retun the neighbors list of all nodes in the graph
func (g *Graph) getNeighbors() [][]Node {
	neighbors := make([][]Node, len(g.NodeList))

	for i, nodeI := range g.NodeList {
		for j, nodeJ := range g.NodeList {

			if i != j && g.isAdjacent(nodeI, nodeJ) {
				neighbors[i] = append(neighbors[i], nodeJ)
			}
		}
	}

	return neighbors
}

/* Return the least distance of all to open vertices list */
func getLeastDistance(distance []float64, verticesName []string, toOpenVertices []string) (float64, int) {

	leastDistance := math.Inf(0)
	leastIndex := 0

	for index, dist := range distance {
		if dist < leastDistance && indexOf(verticesName[index], toOpenVertices) >= 0 {
			leastDistance = dist
			leastIndex = index
		}
	}

	return leastDistance, leastIndex
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

// Dijsktra algorithm
func (g *Graph) Dijsktra(source string) ([]float64, []string) {
	numberVertices := len(g.NodeList)

	/* Created 2D slice */
	var toOpenVertices []string
	neighbors := g.getNeighbors()

	var verticesName []string
	distance := make([]float64, numberVertices)
	previous := make([]string, numberVertices)

	sourceIndex := g.ExistNode(source)

	for key, node := range g.NodeList {
		if key == sourceIndex {
			distance[key] = 0
		} else {
			distance[key] = math.Inf(0)
		}
		previous[key] = "-"

		toOpenVertices = append(toOpenVertices, node.GetName())
		verticesName = append(verticesName, node.GetName())
	}

	for len(toOpenVertices) > 0 {
		_, leastIndex := getLeastDistance(distance, verticesName, toOpenVertices)
		removeIndex := indexOf(verticesName[leastIndex], toOpenVertices)

		toOpenVertices = append(toOpenVertices[:removeIndex], toOpenVertices[removeIndex+1:]...)
		leastNeighbors := neighbors[leastIndex]

		for _, neighbor := range leastNeighbors {
			if index := g.GetEdgeIndex(neighbor.GetName(), g.NodeList[leastIndex].GetName()); index >= 0 &&
				neighbor.GetName() != source {

				neighborIndex := g.ExistNode(neighbor.GetName())

				if (g.EdgeList[index].GetWeight() + distance[leastIndex]) < distance[neighborIndex] {
					distance[neighborIndex] = g.EdgeList[index].GetWeight() + distance[leastIndex]
					previous[neighborIndex] = verticesName[leastIndex]
				}
			}
		}
	}

	return distance, previous
}
