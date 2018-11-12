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

func (g *Graph) GraphvizPNG() {
	g.WriteToFile("output.dot")
	cmd := exec.Command("neato", "-n", "-Tpng", "output.dot", "-o", "./static/graph.png")
	cmd.Run()
}

func (g *Graph) MuxSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		algorithm := r.FormValue("algorithm")
		first_node := r.FormValue("first_node")
		end_node := r.FormValue("end_node")

		if algorithm == "Dijkstra" {
			distance, previous := g.Dijsktra(first_node)
			g.SearchWeight, g.SearchPath = g.DijsktraPath(first_node, end_node, distance, previous)
			g.SearchTable1 = distance
			g.SearchTable2 = previous
			g.ColoringFromPath()
			g.GraphvizPNG()

			// call template render, with graph argument passed.
			rnd.HTML(w, http.StatusOK, "search", g)
		}

		http.Redirect(w, r, "/graph/search", http.StatusSeeOther)

	} else {
		g.SearchWeight = 0.0
		rnd.HTML(w, http.StatusOK, "search", nil)
	}
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

func (graph *Graph) MuxColoring(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if params["algorithm"] == "heuristic" {
		graph.ColoringHeuristic()
	} else if params["algorithm"] == "greedy" {
		graph.Coloring()
	} else {
		graph.ClearColors()
	}

	graph.GraphvizPNG()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (graph *Graph) ValidateNode(node string) bool {
	graph.Errors = make(map[string]string)

	if graph.GetNode(node) != nil {
		graph.Errors["Node"] = "Node " + node + " already exists!"
	}

	return len(graph.Errors) == 0
}

func (graph *Graph) CreateNode(w http.ResponseWriter, r *http.Request) {
	node_name := r.FormValue("vertice_name")
	node_color := r.FormValue("vertice_color")

	if node_color == "" {
		node_color = "white"
	}

	if graph.ValidateNode(node_name) == false {
		rnd.HTML(w, http.StatusOK, "home", graph)
	}

	tmp_node := Node{node_name, node_color, Point{0, 0}}
	graph.AddNode(tmp_node)

	graph.GraphvizPNG()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (graph *Graph) ValidateEdge(edge, first_node, second_node string) bool {
	graph.Errors = make(map[string]string)

	if graph.GetEdge(edge) != nil {
		graph.Errors["Edge"] = "Edge already exists!"
	}

	if graph.GetNode(first_node) == nil {
		graph.Errors["FirstNode"] = "This node don't exists."
	}

	if graph.GetNode(second_node) == nil {
		graph.Errors["SecondNode"] = "This node don't exists."
	}

	return len(graph.Errors) == 0
}

func (graph *Graph) CreateEdge(w http.ResponseWriter, r *http.Request) {
	aresta_name := r.FormValue("aresta_name")
	node_um := r.FormValue("vertice_um")
	node_dois := r.FormValue("vertice_dois")

	if graph.ValidateEdge(aresta_name, node_um, node_dois) == false {
		rnd.HTML(w, http.StatusOK, "home", graph)
	}

	aresta_weight, err := strconv.ParseFloat(r.FormValue("aresta_weight"), 64)
	if err != nil {
		fmt.Println(err)
	}

	first_node := graph.GetNode(node_um)
	second_node := graph.GetNode(node_dois)

	if first_node != nil && second_node != nil {
		edge := Edge{aresta_name, *first_node, *second_node, aresta_weight, ""}

		err = graph.AddEdge(edge)
		if err != nil {
			fmt.Println(err)
		}
	}

	graph.GraphvizPNG()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	var graph = NewGraph()

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

	//graph.ColoringHeuristic()
	// graph.Coloring()

	graph.GraphvizPNG()

	graph.aStar("A", "P")
	//shortestPath, predecessor := graph.Floyd()
	//for i := 0; i < len(shortestPath); i++ {
	//	for j := 0; j < len(shortestPath); j++ {
	//		fmt.Printf("%.0f,", shortestPath[i][j])
	//	}
	//	fmt.Printf("\n")
	//}
	//path := graph.FloydPath(predecessor, "A", "G")

	//fmt.Println("Floyd Shortest Path", path)

	//distance, previous := graph.Dijsktra("A")
	//fmt.Println("Dijkstra Distance", distance)
	//fmt.Println("Dijkstra Predecessor", previous)
	//distanceWeight, dijsktraPath := graph.DijsktraPath("A", "G", distance, previous)
	//fmt.Println("Distancia de A até Q: ", distanceWeight, "Path", dijsktraPath)

	router := mux.NewRouter()
	router.HandleFunc("/", GetGraph)
	router.HandleFunc("/graph/nodes/", graph.CreateNode).Methods("POST")
	router.HandleFunc("/graph/edges/", graph.CreateEdge).Methods("POST")
	router.HandleFunc("/graph/nodes/{name}", graph.GetNodeByName).Methods("GET")
	router.HandleFunc("/graph/color/{algorithm}", graph.MuxColoring).Methods("GET")
	router.HandleFunc("/graph/color/{algorithm}", graph.MuxColoring).Methods("GET")
	router.HandleFunc("/graph/search", graph.MuxSearch).Methods("GET")
	router.HandleFunc("/graph/search", graph.MuxSearch).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(":8000", router)
}
