package main

import "fmt"

func (graph *Graph) testTrabalho() {
	nodeA := Node{"A", "", Point{950, 231}}
	nodeB := Node{"B", "", Point{607, 486}}
	nodeC := Node{"C", "", Point{891, 762}}
	nodeD := Node{"D", "", Point{456, 19}}
	nodeE := Node{"E", "", Point{821, 445}}
	nodeF := Node{"F", "", Point{615, 792}}
	nodeG := Node{"G", "", Point{922, 738}}
	nodeH := Node{"H", "", Point{176, 406}}
	nodeI := Node{"I", "", Point{935, 917}}
	nodeJ := Node{"J", "", Point{410, 894}}
	nodeK := Node{"K", "", Point{58, 353}}
	nodeL := Node{"L", "", Point{813, 010}}
	nodeM := Node{"M", "", Point{139, 203}}
	nodeN := Node{"N", "", Point{199, 604}}
	nodeO := Node{"O", "", Point{272, 199}}
	nodeP := Node{"P", "", Point{015, 747}}
	nodeQ := Node{"Q", "", Point{445, 932}}
	nodeR := Node{"R", "", Point{466, 419}}
	nodeS := Node{"S", "", Point{846, 525}}
	nodeT := Node{"T", "", Point{203, 672}}

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
