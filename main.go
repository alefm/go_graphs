package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// GetGraph - this is a handler to / requisition
func GetGraph(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entrei graph")
	return
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

	node1 := Node{"1", ""}
	node2 := Node{"2", ""}
	node3 := Node{"3", ""}
	node4 := Node{"4", ""}

	graph.AddNode(node1)
	graph.AddNode(node2)
	graph.AddNode(node3)
	graph.AddNode(node4)

	edge1 := Edge{"a", node2, node1, 4, ""}
	edge2 := Edge{"b", node1, node3, -2, ""}
	edge3 := Edge{"c", node3, node4, 2, ""}
	edge4 := Edge{"d", node4, node2, -1, ""}
	edge5 := Edge{"e", node2, node3, 3, ""}

	err := graph.AddEdge(edge1)
	if err != nil {
		fmt.Println(err)
	}

	err2 := graph.AddEdge(edge2)
	if err2 != nil {
		fmt.Println(err2)
	}

	err3 := graph.AddEdge(edge3)
	if err3 != nil {
		fmt.Println(err3)
	}

	err4 := graph.AddEdge(edge4)
	if err4 != nil {
		fmt.Println(err4)
	}

	err5 := graph.AddEdge(edge5)
	if err4 != nil {
		fmt.Println(err5)
	}

	graph.WriteToFile("output.dot")
	_, predecessor := graph.FloydAlgorithm()
	path := graph.FloydPath(predecessor, node1.Name, node4.Name)

	fmt.Println(path)

	router := mux.NewRouter()
	router.HandleFunc("/", GetGraph)
	router.HandleFunc("/graph/nodes/{name}", graph.GetNodeByName).Methods("GET")

	http.ListenAndServe(":8001", router)
}
