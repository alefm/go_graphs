package main

import (
	"fmt"
	"math"
)

type distanceHeuristic struct {
	source      Node
	destination Node
	distance    float64
}

func calculateDistance(source Node, destination Node) float64 {
	sumX := math.Abs(float64(source.GraphPoint.X - destination.GraphPoint.X))
	sumY := math.Abs(float64(source.GraphPoint.Y - destination.GraphPoint.Y))
	return sumX + sumY
}

func getLowestDistance(list []distanceHeuristic) (*distanceHeuristic, int) {
	if len(list) <= 0 {
		return nil, -1
	}

	lowestDistance := list[0]
	lowestIdx := 0
	for i, distance := range list {
		if distance.distance < lowestDistance.distance {
			lowestDistance = distance
			lowestIdx = i
		}
	}

	return &lowestDistance, lowestIdx
}

func isInNodeList(list []Node, name string) bool {
	for _, node := range list {
		if name == node.GetName() {
			return true
		}
	}
	return false
}

func isInHeuristicList(list []distanceHeuristic, name string) (bool, int) {
	for i, dist := range list {
		if name == dist.destination.GetName() {
			return true, i
		}
	}
	return false, -1
}

func (g *Graph) distanceBetween(source string, end string) float64 {
	for _, edge := range g.EdgeList {
		if (edge.begin.Name == source && edge.end.Name == end) ||
			(edge.begin.Name == end && edge.end.Name == source) {
			return edge.weight
		}
	}
	return 0
}

func (g *Graph) aStar(source string, end string) {
	sourceNode := *g.GetNode(source)
	endNode := *g.GetNode(end)
	var distanceList []distanceHeuristic
	var openList []distanceHeuristic
	var closedList []Node
	var cameFrom []string
	neighbors := g.getNeighbors()

	//gscore -> para cada nodo valor do node inicial até o nodo corrente
	//fscore -> para cada nodo o custo total do nodo inicial até o final passando pelo nodo corrente
	//https://en.wikipedia.org/wiki/A*_search_algorithm

	// Generate distance list of source to all points
	for _, node := range g.NodeList {
		if sourceNode.GetName() != node.GetName() {
			distance := calculateDistance(sourceNode, node)
			distanceList = append(distanceList, distanceHeuristic{sourceNode, node, distance})
		} else {
			distSource := distanceHeuristic{sourceNode, node, 0}
			distanceList = append(distanceList, distSource)

			// Put source node in openList
			openList = append(openList, distSource)
		}
	}

	for len(openList) > 0 {
		lowestD, lowestIdx := getLowestDistance(openList)
		q := lowestD.destination

		if q.GetName() == end {
			cameFrom = append(cameFrom, q.GetName())
			fmt.Println(cameFrom)
		}

		// Remove lowest node distance from openList
		openList = append(openList[:lowestIdx], openList[lowestIdx+1:]...)

		// Put the lowest node distance in closedList
		closedList = append(closedList, q)

		// Get all neighbors from this node
		qNeighbors := neighbors[g.ExistNode(q.GetName())]

		for _, n := range qNeighbors {
			if isInNodeList(closedList, n.GetName()) {
				continue // Ignore neighbor in closedList
			}

			_, tentativeIdx := isInHeuristicList(distanceList, n.GetName())

			// Sum calculated distance with real distance
			tentativeDistance := distanceList[tentativeIdx].distance + g.distanceBetween(q.GetName(), n.GetName())

			// Discover a new node
			if present, _ := isInHeuristicList(openList, n.GetName()); !present {
				openList = append(openList, distanceList[tentativeIdx])
			} else if tentativeDistance >= distanceList[tentativeIdx].distance {
				continue // Ignore new distance, the older is shortest
			}

			cameFrom = append(cameFrom, q.GetName())
			distanceList[tentativeIdx].distance = tentativeDistance + calculateDistance(n, endNode)
		}
	}
}
