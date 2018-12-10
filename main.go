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

func (g *Graph) GraphvizPNG(filename string) {
	g.WriteToFile(filename + ".dot")
	cmd := exec.Command("neato", "-n", "-Tpng", filename+".dot", "-o", "./static/"+filename+".png")
	cmd.Run()
}

func (g *Graph) MuxSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		algorithm := r.FormValue("algorithm")
		first_node := r.FormValue("first_node")
		end_node := r.FormValue("end_node")

		g.SearchPath = nil
		g.SearchWeight = 0.0
		g.SearchTable1 = nil
		g.SearchTable2 = nil
		g.SearchTable3 = nil
		g.SearchTable4 = nil

		if algorithm == "Dijkstra" {
			distance, previous := g.Dijsktra(first_node)
			g.SearchWeight, g.SearchPath = g.DijsktraPath(first_node, end_node, distance, previous)
			g.SearchTable1 = distance
			g.SearchTable2 = previous
			g.ColoringFromPath()
			g.GraphvizPNG("graph")

			// call template render, with graph argument passed.
			rnd.HTML(w, http.StatusOK, "search", g)
		} else if algorithm == "Floyd" {

			g.SearchTable3, g.SearchTable4 = g.Floyd()
			g.SearchPath = g.FloydPath(g.SearchTable4, first_node, end_node)

			fmt.Println("**************************** FLOYD ****************************")
			fmt.Println("Distance Matrix:")

			for _, line := range g.SearchTable3 {
				for _, distance := range line {
					fmt.Printf("%.0f, ", distance)
				}
				fmt.Printf("\n")
			}

			fmt.Println("Predecessor Matrix:")
			for _, line := range g.SearchTable4 {
				for _, predecessor := range line {
					if predecessor == "" {
						fmt.Printf("-, ")
					} else {
						fmt.Printf("%s, ", predecessor)
					}
				}
				fmt.Printf("\n")
			}

			fmt.Println("***************************************************************")
			g.ColoringFromPath()
			g.GraphvizPNG("graph")

			// call template render, with graph argument passed.
			rnd.HTML(w, http.StatusOK, "search", g)
		} else if algorithm == "A*" {

			g.SearchPath, g.SearchWeight = g.aStar(first_node, end_node)
			g.ColoringFromPath()
			g.GraphvizPNG("graph")

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

	graph.GraphvizPNG("graph")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (graph *Graph) ValidateNode(node string, x, y float64) bool {
	graph.Errors = make(map[string]string)

	if graph.GetNode(node) != nil {
		graph.Errors["Node"] = "Node " + node + " already exists!"
	}

	if x < 0.0 || x > 1000.0 {
		graph.Errors["X"] = "X: 0~1000"
	}

	if y < 0.0 || y > 1000.0 {
		graph.Errors["Y"] = "Y: 0~1000"
	}

	return len(graph.Errors) == 0
}

func (graph *Graph) CreateNode(w http.ResponseWriter, r *http.Request) {
	node_name := r.FormValue("node_name")
	node_color := r.FormValue("node_color")
	node_x, x_err := strconv.ParseFloat(r.FormValue("node_x"), 64)
	node_y, y_err := strconv.ParseFloat(r.FormValue("node_y"), 64)

	if x_err != nil || y_err != nil {
		fmt.Println(y_err)
	}

	tmp_node := Node{node_name, node_color, Point{node_x, node_y}}
	graph.AddNode(tmp_node)

	graph.GraphvizPNG("graph")

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
	edge_name := r.FormValue("edge_name")
	edge_weight, err := strconv.ParseFloat(r.FormValue("edge_weight"), 64)
	if err != nil {
		fmt.Println(err)
	}
	node_one := r.FormValue("node_one")
	node_two := r.FormValue("node_two")

	if graph.ValidateEdge(edge_name, node_one, node_two) == false {
		rnd.HTML(w, http.StatusOK, "home", graph)
	}

	first_node := graph.GetNode(node_one)
	second_node := graph.GetNode(node_two)

	if first_node != nil && second_node != nil {
		edge := Edge{edge_name, *first_node, *second_node, edge_weight, ""}

		err = graph.AddEdge(edge)
		if err != nil {
			fmt.Println(err)
		}
	}

	graph.GraphvizPNG("graph")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Genetic Algorithm
func (graph *Graph) GeneticIndex(w http.ResponseWriter, r *http.Request) {
	graph.GraphvizPNG("graph")

	//rnd.HTML(w, http.StatusOK, "genetic", graph )
	rnd.HTML(w, http.StatusOK, "genetic", graph.frontend.genetic)
}

func (graph *Graph) GeneticExperiment(w http.ResponseWriter, r *http.Request) {

	node_begin := r.FormValue("node_begin")
	population, err := strconv.Atoi(r.FormValue("population"))
	stop, err := strconv.Atoi(r.FormValue("stop"))
	crossover, err := strconv.ParseFloat(r.FormValue("crossover"), 64)
	mutation, err := strconv.ParseFloat(r.FormValue("mutation"), 64)

	if err != nil {
		fmt.Println(err)
	}

	solution_path, _ := graph.geneticAlgorithm(node_begin, int(stop), crossover, mutation, int(population))

	graph.SearchPath = solution_path

	graph.GraphvizPNG("graph")

	var tsp_graph = NewGraph()

	var edge Edge
	for _, letter := range solution_path {
		node := graph.GetNode(letter)
		tsp_graph.AddNode(*node)
	}
	for idx, _ := range solution_path {

		if idx+1 > len(solution_path)-1 {
			node_one := graph.GetNode(solution_path[idx])
			node_two := graph.GetNode(solution_path[0])
			edge = Edge{node_one.Name + node_two.Name, *node_one, *node_two, 0, ""}
			tsp_graph.AddEdge(edge)
		} else {
			node_one := graph.GetNode(solution_path[idx])
			node_two := graph.GetNode(solution_path[idx+1])
			edge = Edge{node_one.Name + node_two.Name, *node_one, *node_two, 0, ""}
			tsp_graph.AddEdge(edge)
		}
	}

	tsp_graph.GraphvizPNG("genetic")

	http.Redirect(w, r, "/graph/genetic", http.StatusSeeOther)
}

func main() {
	var graph = NewGraph()

	graph.testTrabalhoM3()
	graph.geneticAlgorithm("H", 10, 60, 1, 100)
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

	graph.GraphvizPNG("graph")

	router := mux.NewRouter()
	router.HandleFunc("/", GetGraph)
	router.HandleFunc("/graph/nodes", graph.CreateNode).Methods("POST")
	router.HandleFunc("/graph/edges", graph.CreateEdge).Methods("POST")
	router.HandleFunc("/graph/nodes/{name}", graph.GetNodeByName).Methods("GET")
	router.HandleFunc("/graph/color/{algorithm}", graph.MuxColoring).Methods("GET")
	router.HandleFunc("/graph/color/{algorithm}", graph.MuxColoring).Methods("GET")
	router.HandleFunc("/graph/search", graph.MuxSearch).Methods("GET")
	router.HandleFunc("/graph/search", graph.MuxSearch).Methods("POST")
	router.HandleFunc("/graph/genetic", graph.GeneticIndex).Methods("GET")
	router.HandleFunc("/graph/genetic", graph.GeneticExperiment).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(":8000", router)
}
