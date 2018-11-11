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

func (g *Graph) aStar(source string, end string) {
	sourceNode := *g.GetNode(source)
	var distanceList []distanceHeuristic
	var openList []distanceHeuristic
	var closedList []Node
	var cameFrom []string
	neighbors := g.getNeighbors()

	// Generate distance list of source to all points
	for _, node := range g.NodeList {
		if sourceNode.GetName() != node.GetName() {
			distance := calculateDistance(sourceNode, node)
			distanceList = append(distanceList, distanceHeuristic{sourceNode, node, distance})
		} else {
			distSource := distanceHeuristic{sourceNode, node, 0}
			distanceList = append(distanceList, distSource)
			openList = append(openList, distSource)
		}
	}

	for len(openList) > 0 {
		lowestD, lowestIdx := getLowestDistance(openList)
		q := lowestD.destination

		if q.GetName() == end {
			fmt.Println(cameFrom)
		}

		openList = append(openList[:lowestIdx], openList[lowestIdx+1:]...)
		closedList = append(closedList, q)
		qNeighbors := neighbors[g.ExistNode(q.GetName())]

		for _, n := range qNeighbors {
			if isInNodeList(closedList, n.GetName()) {
				continue // Ignore neighbor in closedList
			}

			tentativeDistance := distanceList[g.ExistNode(n.GetName())].distance + g.EdgeList[g.GetEdgeIndex(q.GetName(), n.GetName())].weight //lowestD.distance + calculateDistance(q, n)

			if present, _ := isInHeuristicList(openList, n.GetName()); !present {
				_, idx := isInHeuristicList(distanceList, n.GetName())
				openList = append(openList, distanceList[idx])
			} else if tentativeDistance >= distanceList[g.ExistNode(n.GetName())].distance {
				continue
			}

			cameFrom = append(cameFrom, q.GetName())
			distanceList[g.ExistNode(n.GetName())].distance = tentativeDistance
		}
	}
}
