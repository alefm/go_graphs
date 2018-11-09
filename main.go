package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./templates/*.html",
	}

	rnd = renderer.New(opts)
}

// GetGraph - this is a handler to / requisition
func GetGraph(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "home", nil)
}

// GetNodeByName - this is a handler to /nodes/{name} requisition
func (graph *Graph) GetNodeByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range graph.NodeList {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(&item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Node{})
}

func main() {

	graph := NewGraph()

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

	edgeList = append(edgeList, Edge{"", nodeC, nodeG, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeQ, nodeI, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeI, nodeC, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeI, nodeG, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeF, nodeJ, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeQ, nodeJ, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeC, nodeF, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeJ, nodeP, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeJ, nodeT, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeJ, nodeB, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeF, nodeB, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeF, nodeS, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeP, nodeT, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeT, nodeB, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeG, nodeS, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeB, nodeS, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeT, nodeN, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeP, nodeN, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeN, nodeR, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeB, nodeE, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeG, nodeA, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeS, nodeA, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeE, nodeA, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeB, nodeA, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeN, nodeH, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeP, nodeH, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeH, nodeO, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeO, nodeR, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeR, nodeB, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeA, nodeL, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeR, nodeL, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeR, nodeD, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeM, nodeO, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeK, nodeM, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeD, nodeL, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeP, nodeK, 4, ""})
	edgeList = append(edgeList, Edge{"", nodeM, nodeD, 4, ""})

	for _, edge := range edgeList {
		err := graph.AddEdge(edge)
		if err != nil {
			fmt.Println(err)
		}
	}

	// graph.ColoringHeuristic()
	graph.Coloring()
	graph.aStar("A")

	graph.WriteToFile("output.dot")
	cmd := exec.Command("neato", "-n", "-Tpng", "output.dot", "-o", "./static/graph.png")
	cmd.Run()

	// _, predecessor := graph.Floyd()
	// path := graph.FloydPath(predecessor, nodeA.Name, nodeQ.Name)

	// fmt.Println("Floyd Shortest Path", path)

	distance, previous := graph.Dijsktra("A")
	fmt.Println("Dijkstra Distance", distance)
	fmt.Println("Dijkstra Predecessor", previous)
	distanceWeight, dijsktraPath := graph.DijsktraPath("A", "Q", distance, previous)
	fmt.Println("Distancia de A até Q: ", distanceWeight, "Path", dijsktraPath)

	router := mux.NewRouter()
	router.HandleFunc("/", GetGraph)
	router.HandleFunc("/graph/nodes/{name}", graph.GetNodeByName).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(":8000", router)
}
