package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strconv" 

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

func (graph *Graph) ValidateNode(node string) bool {
	graph.Errors = make(map[string]string)
	
	if graph.GetNode(node) != nil {
		graph.Errors["Node"] = "Node already exists!"
	}
	
	return len(graph.Errors) == 0
}

func (graph *Graph) CreateNode(w http.ResponseWriter, r *http.Request) {
	node_name := r.FormValue("vertice_name")
	node_weight := r.FormValue("vertice_weight")

	if graph.ValidateNode(node_name) == false {
		rnd.HTML(w, http.StatusOK, "home", graph)
	}

	tmp_node := Node{node_name, node_weight}
	graph.AddNode(tmp_node)

	graph.WriteToFile("output.dot")
	cmd := exec.Command("dot", "-Tpng", "output.dot", "-o", "./static/graph.png")
	cmd.Run()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (graph *Graph) ValidateEdge(edge string) bool {
	graph.Errors = make(map[string]string)
	
	if graph.GetEdge(edge) != nil {
		graph.Errors["Edge"] = "Edge already exists!"
	}
	
	return len(graph.Errors) == 0
}

func (graph *Graph) ValidateEdgeNode(first_node, second_node string) bool {
	graph.Errors = make(map[string]string)
	
	if graph.GetNode(first_none) == nil {
		graph.Errors["FirstNode"] = "This node don't exists."
	}

	if graph.GetNode(second_node) == nil {
		graph.Errors["SecondNode"] = "This node don't exists."
	}
	
	return len(graph.Errors) == 0
}

func (graph *Graph) CreateEdge(w http.ResponseWriter, r *http.Request) {
	aresta_name := r.FormValue("aresta_name")
	aresta_weight, err := strconv.ParseFloat(r.FormValue("aresta_weight"), 64)
	if err != nil {
		fmt.Println(err)
	}

	node_um := r.FormValue("vertice_um")
	node_dois := r.FormValue("vertice_dois")
	
	first_node := graph.GetNode(node_um)
	second_node := graph.GetNode(node_dois)

	edge := Edge{aresta_name, *first_node, *second_node , aresta_weight, ""}

	err = graph.AddEdge(edge)
	if err != nil {
		fmt.Println(err)
	}

	graph.WriteToFile("output.dot")
	cmd := exec.Command("dot", "-Tpng", "output.dot", "-o", "./static/graph.png")
	cmd.Run()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	var graph = NewGraph()


	/**
	node1 := Node{"1", ""}
	node2 := Node{"2", ""}
	node3 := Node{"3", ""}
	node4 := Node{"4", ""}
	node5 := Node{"5", ""}
	node6 := Node{"6", ""}

	graph.AddNode(node1)
	graph.AddNode(node2)
	graph.AddNode(node3)
	graph.AddNode(node4)
	graph.AddNode(node5)
	graph.AddNode(node6)

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

	*/
	graph.WriteToFile("output.dot")
	cmd := exec.Command("dot", "-Tpng", "output.dot", "-o", "./static/graph.png")
	cmd.Run()

	/**
	_, predecessor := graph.FloydAlgorithm()
	path := graph.FloydPath(predecessor, node1.Name, node4.Name)

	fmt.Println(path)
	*/

	router := mux.NewRouter()
	router.HandleFunc("/", GetGraph)
	router.HandleFunc("/graph/nodes/", graph.CreateNode).Methods("POST")
	router.HandleFunc("/graph/edges/", graph.CreateEdge).Methods("POST")
	router.HandleFunc("/graph/nodes/{name}", graph.GetNodeByName).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(":8000", router)
}
