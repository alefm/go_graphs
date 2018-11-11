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

	graph.testTrabalho()
	/*node0 := Node{"0", "", Point{4.10, 8.94}}
	node1 := Node{"1", "", Point{9.50, 2.31}}
	node2 := Node{"2", "", Point{6.07, 4.86}}
	node3 := Node{"3", "", Point{8.91, 7.62}}
	node4 := Node{"4", "", Point{4.56, 0.19}}
	node5 := Node{"5", "", Point{8.21, 4.45}}
	node6 := Node{"6", "", Point{6.15, 7.92}}
	node7 := Node{"7", "", Point{9.22, 7.38}}
	node8 := Node{"8", "", Point{1.76, 4.06}}
	node9 := Node{"9", "", Point{9.35, 9.17}}

	graph.AddNode(node0)
	graph.AddNode(node1)
	graph.AddNode(node2)
	graph.AddNode(node3)
	graph.AddNode(node4)
	graph.AddNode(node5)
	graph.AddNode(node6)
	graph.AddNode(node7)
	graph.AddNode(node8)
	graph.AddNode(node9)

	var edgeList []Edge

	edgeList = append(edgeList, Edge{"", node4, node8, 6, ""})
	edgeList = append(edgeList, Edge{"", node5, node4, 4, ""})
	edgeList = append(edgeList, Edge{"", node5, node8, 12, ""})
	edgeList = append(edgeList, Edge{"", node8, node9, 13, ""})
	edgeList = append(edgeList, Edge{"", node8, node6, 16, ""})
	edgeList = append(edgeList, Edge{"", node6, node9, 6, ""})
	edgeList = append(edgeList, Edge{"", node3, node5, 6, ""})
	edgeList = append(edgeList, Edge{"", node3, node6, 3, ""})
	edgeList = append(edgeList, Edge{"", node1, node4, 14, ""})
	edgeList = append(edgeList, Edge{"", node1, node3, 11, ""})
	edgeList = append(edgeList, Edge{"", node0, node1, 6, ""})
	edgeList = append(edgeList, Edge{"", node1, node2, 12, ""})
	edgeList = append(edgeList, Edge{"", node0, node2, 10, ""})
	edgeList = append(edgeList, Edge{"", node2, node3, 12, ""})
	edgeList = append(edgeList, Edge{"", node2, node6, 8, ""})
	edgeList = append(edgeList, Edge{"", node2, node7, 16, ""})
	edgeList = append(edgeList, Edge{"", node7, node9, 8, ""})

	for _, edge := range edgeList {
		err := graph.AddEdge(edge)
		if err != nil {
			fmt.Println(err)
		}
	}*/

	graph.ColoringHeuristic()
	// graph.Coloring()

	graph.WriteToFile("output.dot")
	cmd := exec.Command("neato", "-n", "-Tpng", "output.dot", "-o", "./static/graph.png")
	cmd.Run()

	// graph.aStar("A")
	_, predecessor := graph.Floyd()

	path := graph.FloydPath(predecessor, "A", "P")

	fmt.Println("Floyd Shortest Path", path)

	distance, previous := graph.Dijsktra("A")
	fmt.Println("Dijkstra Distance", distance)
	fmt.Println("Dijkstra Predecessor", previous)
	distanceWeight, dijsktraPath := graph.DijsktraPath("A", "P", distance, previous)
	fmt.Println("Distancia de A até P: ", distanceWeight, "Path", dijsktraPath)

	router := mux.NewRouter()
	router.HandleFunc("/", GetGraph)
	router.HandleFunc("/graph/nodes/{name}", graph.GetNodeByName).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(":8000", router)
}
