package main

import (
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

func reconstructPath(cameFrom map[string]string, current string) []string {
	var totalPath []string
	_, ok := cameFrom[current]
	currentAux := cameFrom[current]

	totalPath = append([]string{current}, totalPath...)
	totalPath = append([]string{currentAux}, totalPath...)

	for ok {
		currentAux = cameFrom[currentAux]
		totalPath = append([]string{currentAux}, totalPath...)
		_, ok = cameFrom[currentAux]
	}

	return totalPath
}

func (g *Graph) aStar(source string, end string) ([]string, float64) {
	sourceNode := *g.GetNode(source)
	endNode := *g.GetNode(end)
	// distance between source until node
	gScore := make(map[string]distanceHeuristic)
	// distance between source to end passing by node
	fScore := make(map[string]float64)

	var distanceList []distanceHeuristic
	var openList []distanceHeuristic
	var closedList []Node

	cameFrom := make(map[string]string)
	neighbors := g.getNeighbors()

	// Generate distance list of source to all points
	for _, node := range g.NodeList {
		if sourceNode.GetName() != node.GetName() {
			distance := calculateDistance(sourceNode, node)

			distanceToEnd := calculateDistance(node, endNode)
			distanceToEnd = distanceToEnd + distance

			distanceH := distanceHeuristic{sourceNode, node, distance}
			gScore[node.GetName()] = distanceH
			fScore[node.GetName()] = distanceToEnd

			distanceList = append(distanceList, distanceH)
		} else {
			distSource := distanceHeuristic{sourceNode, node, 0}
			gScore[node.GetName()] = distSource
			fScore[node.GetName()] = calculateDistance(sourceNode, endNode)

			// Put source node in openList
			openList = append(openList, distSource)

			distanceList = append(distanceList, distSource)
		}
	}

	for len(openList) > 0 {
		lowestD, lowestIdx := getLowestDistance(openList)
		q := lowestD.destination

		if q.GetName() == end {
			return reconstructPath(cameFrom, q.GetName()), gScore[endNode.GetName()].distance
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
			tentativeDistance := gScore[q.GetName()].distance + calculateDistance(q, n)

			// Discover a new node
			if present, _ := isInHeuristicList(openList, n.GetName()); !present {
				openList = append(openList, distanceList[tentativeIdx])
			} else if tentativeDistance >= gScore[n.GetName()].distance {
				continue // Ignore new distance, the older is shortest
			}

			cameFrom[n.GetName()] = q.GetName()
			aux := gScore[n.GetName()]
			aux.distance = tentativeDistance

			gScore[n.GetName()] = aux
			fScore[n.GetName()] = gScore[n.GetName()].distance + calculateDistance(n, endNode)
		}
	}

	return make([]string, 0), 0.0
}
