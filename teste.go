package main

import "fmt"

func (graph *Graph) testTrabalho() {
	nodeA := Node{"A", "", Point{9.50, 2.31}}
	nodeB := Node{"B", "", Point{6.07, 4.86}}
	nodeC := Node{"C", "", Point{8.91, 7.62}}
	nodeD := Node{"D", "", Point{4.56, 0.19}}
	nodeE := Node{"E", "", Point{8.21, 4.45}}
	nodeF := Node{"F", "", Point{6.15, 7.92}}
	nodeG := Node{"G", "", Point{9.22, 7.38}}
	nodeH := Node{"H", "", Point{1.76, 4.06}}
	nodeI := Node{"I", "", Point{9.35, 9.17}}
	nodeJ := Node{"J", "", Point{4.10, 8.94}}
	nodeK := Node{"K", "", Point{0.58, 3.53}}
	nodeL := Node{"L", "", Point{8.13, 0.10}}
	nodeM := Node{"M", "", Point{1.39, 2.03}}
	nodeN := Node{"N", "", Point{1.99, 6.04}}
	nodeO := Node{"O", "", Point{2.72, 1.99}}
	nodeP := Node{"P", "", Point{0.15, 7.47}}
	nodeQ := Node{"Q", "", Point{4.45, 9.32}}
	nodeR := Node{"R", "", Point{4.66, 4.19}}
	nodeS := Node{"S", "", Point{8.46, 5.25}}
	nodeT := Node{"T", "", Point{2.03, 6.72}}

	graph.AddNode(nodeA)
	graph.AddNode(nodeB)
	graph.AddNode(nodeC)
	graph.AddNode(nodeD)
	graph.AddNode(nodeE)
	graph.AddNode(nodeF)
	graph.AddNode(nodeG)
	graph.AddNode(nodeH)
	graph.AddNode(nodeI)
	graph.AddNode(nodeJ)
	graph.AddNode(nodeK)
	graph.AddNode(nodeL)
	graph.AddNode(nodeM)
	graph.AddNode(nodeN)
	graph.AddNode(nodeO)
	graph.AddNode(nodeP)
	graph.AddNode(nodeQ)
	graph.AddNode(nodeR)
	graph.AddNode(nodeS)
	graph.AddNode(nodeT)

	var edgeList []Edge

	edgeList = append(edgeList, Edge{"", nodeC, nodeG, 40, ""})
	edgeList = append(edgeList, Edge{"", nodeQ, nodeI, 90, ""})
	edgeList = append(edgeList, Edge{"", nodeI, nodeC, 46, ""})
	edgeList = append(edgeList, Edge{"", nodeI, nodeG, 21, ""})
	edgeList = append(edgeList, Edge{"", nodeF, nodeJ, 78, ""})
	edgeList = append(edgeList, Edge{"", nodeQ, nodeJ, 99, ""})
	edgeList = append(edgeList, Edge{"", nodeC, nodeF, 12, ""})
	edgeList = append(edgeList, Edge{"", nodeJ, nodeP, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeJ, nodeT, 54, ""})
	edgeList = append(edgeList, Edge{"", nodeJ, nodeB, 123, ""})
	edgeList = append(edgeList, Edge{"", nodeF, nodeB, 90, ""})
	edgeList = append(edgeList, Edge{"", nodeF, nodeS, 87, ""})
	edgeList = append(edgeList, Edge{"", nodeP, nodeT, 52, ""})
	edgeList = append(edgeList, Edge{"", nodeT, nodeB, 31, ""})
	edgeList = append(edgeList, Edge{"", nodeG, nodeS, 21, ""})
	edgeList = append(edgeList, Edge{"", nodeB, nodeS, 11, ""})
	edgeList = append(edgeList, Edge{"", nodeT, nodeN, 92, ""})
	edgeList = append(edgeList, Edge{"", nodeP, nodeN, 18, ""})
	edgeList = append(edgeList, Edge{"", nodeN, nodeR, 43, ""})
	edgeList = append(edgeList, Edge{"", nodeB, nodeE, 34, ""})
	edgeList = append(edgeList, Edge{"", nodeG, nodeA, 81, ""})
	edgeList = append(edgeList, Edge{"", nodeS, nodeA, 87, ""})
	edgeList = append(edgeList, Edge{"", nodeE, nodeA, 66, ""})
	edgeList = append(edgeList, Edge{"", nodeB, nodeA, 44, ""})
	edgeList = append(edgeList, Edge{"", nodeN, nodeH, 22, ""})
	edgeList = append(edgeList, Edge{"", nodeP, nodeH, 83, ""})
	edgeList = append(edgeList, Edge{"", nodeH, nodeO, 12, ""})
	edgeList = append(edgeList, Edge{"", nodeO, nodeR, 89, ""})
	edgeList = append(edgeList, Edge{"", nodeR, nodeB, 76, ""})
	edgeList = append(edgeList, Edge{"", nodeA, nodeL, 14, ""})
	edgeList = append(edgeList, Edge{"", nodeR, nodeL, 84, ""})
	edgeList = append(edgeList, Edge{"", nodeR, nodeD, 55, ""})
	edgeList = append(edgeList, Edge{"", nodeM, nodeO, 21, ""})
	edgeList = append(edgeList, Edge{"", nodeK, nodeM, 14, ""})
	edgeList = append(edgeList, Edge{"", nodeD, nodeL, 62, ""})
	edgeList = append(edgeList, Edge{"", nodeP, nodeK, 71, ""})
	edgeList = append(edgeList, Edge{"", nodeM, nodeD, 77, ""})

	for _, edge := range edgeList {
		err := graph.AddEdge(edge)
		if err != nil {
			fmt.Println(err)
		}
	}
}
