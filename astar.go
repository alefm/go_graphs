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

func (g *Graph) aStar(source string) {
	sourceNode := *g.GetNode(source)
	var distanceList []distanceHeuristic

	// Generate distance list of source to all points
	for _, node := range g.NodeList {
		if sourceNode.GetName() != node.GetName() {
			distance := calculateDistance(sourceNode, node)
			distanceList = append(distanceList, distanceHeuristic{sourceNode, node, distance})
		} else {
			distanceList = append(distanceList, distanceHeuristic{sourceNode, node, 0})
		}
	}

	for _, distanceH := range distanceList {
		fmt.Printf("Distance %s -> %s = %.2f\n", distanceH.source.Name, distanceH.destination.Name, distanceH.distance)
	}
}
