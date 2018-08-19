package main

import (
	"fmt"
	"time"
)

type Node struct {
	// position_x float64
	// position_y float64
	name            string
	next            *Node
	connection_list *NodeList
	next_connection *Node
}

type NodeList struct {
	length int
	start  *Node
}

func printConnectionList(node Node) {
	if node.connection_list.length == 0 {
		fmt.Printf("Não há nenhuma Conexão!\n")
	} else {
		currentNode := node.connection_list.start
		for currentNode != nil {
			fmt.Printf("%s ", currentNode.name)
			currentNode = currentNode.next_connection
		}

		fmt.Printf("\n")
	}
}

func printNodeList(list NodeList) {
	fmt.Printf("Lista de nodos: \n")
	if list.length == 0 {
		fmt.Printf("Não há nenhum nodo!\n")
	} else {
		currentNode := list.start
		for currentNode != nil {
			fmt.Printf("Nodo %s: ", currentNode.name)
			printConnectionList(*currentNode)
			currentNode = currentNode.next
		}
	}
}

func (list *NodeList) Append(newNode *Node) {
	if list.length == 0 {
		list.start = newNode
	} else {
		currentNode := list.start
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
	list.length++
}

func (node *Node) AppendConnection(newConnection *Node) {
	if node.connection_list.length == 0 {
		node.connection_list.start = newConnection
	} else {
		currentNode := node.connection_list.start
		for currentNode.next_connection != nil {
			currentNode = currentNode.next_connection
		}
		currentNode.next_connection = newConnection
	}
	node.connection_list.length++
}

func (list *NodeList) RemoveNode(node *Node) *Node {
	var previousNode *Node
	previousNode = list.start
	currentNode := list.start

	if list.length == 0 {
		fmt.Printf("Lista Vazia!\n")
		return nil
	}

	//TODO: remoção por node ID
	for currentNode.name != node.name {
		if currentNode.next == nil {
			fmt.Println("Nodo não encontrado")
		}

		previousNode = currentNode
		currentNode = currentNode.next
	}

	previousNode.next = currentNode.next

	if currentNode.connection_list.length > 0 {
		currentConnected := currentNode.connection_list.start
		for currentConnected != nil {
			RemoveConnection(currentNode, currentConnected)
			currentConnected = currentConnected.next_connection
		}
	}

	currentNode.next = nil

	list.length--

	return currentNode
}

func (node *Node) RemoveNodeConnection(node_removed *Node) *Node {
	if node.connection_list.length == 0 {
		fmt.Printf("Lista Vazia!\n")
		return nil
	}

	var previousNode *Node
	previousNode = node.connection_list.start
	currentNode := node.connection_list.start

	//TODO: remoção por node ID
	for currentNode.name != node_removed.name {
		if currentNode.next_connection == nil {
			fmt.Println("Nodo não encontrado %s %s", currentNode.name, node_removed.name)
		}

		previousNode = currentNode
		currentNode = currentNode.next_connection
	}

	previousNode.next_connection = currentNode.next_connection
	currentNode.next_connection = nil

	node.connection_list.length--

	return currentNode
}

func RemoveConnection(firstNode *Node, secondNode *Node) {
	firstNode.RemoveNodeConnection(secondNode)
	secondNode.RemoveNodeConnection(firstNode)
}

func makeConnection(firstNode *Node, secondNode *Node) {
	firstNode.AppendConnection(secondNode)
	secondNode.AppendConnection(firstNode)
}

func main() {
	node_list := &NodeList{}

	nodeA := Node{
		name:            "A",
		connection_list: &NodeList{},
	}

	nodeB := Node{
		name:            "B",
		connection_list: &NodeList{},
	}

	nodeC := Node{
		name:            "C",
		connection_list: &NodeList{},
	}

	nodeD := Node{
		name:            "D",
		connection_list: &NodeList{},
	}

	node_list.Append(&nodeA)
	node_list.Append(&nodeB)
	node_list.Append(&nodeC)
	node_list.Append(&nodeD)

	makeConnection(&nodeA, &nodeB)
	makeConnection(&nodeA, &nodeC)
	makeConnection(&nodeA, &nodeD)

	printNodeList(*node_list)

	fmt.Printf("Removendo Nodo\n")
	time.Sleep(100 * time.Millisecond)
	nodo_removido := node_list.RemoveNode(&nodeC)

	printNodeList(*node_list)
	if nodo_removido != nil {
		fmt.Printf("Nodo removido %s\n", nodo_removido.name)
	}

	fmt.Printf("Removendo Conexao\n")
	time.Sleep(100 * time.Millisecond)

	RemoveConnection(&nodeA, &nodeD)
	printNodeList(*node_list)
}

/*func main() {

	rand.Seed(time.Now().UnixNano())

	dc := gg.NewContext(screen_width, screen_height)

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(0, 0, 0)
	x, y := random_position()
	dc.DrawCircle(x, y, 20)
	dc.Fill()

	dc.SetRGB(1, 1, 1)

	if err := dc.LoadFontFace("/Library/Fonts/Arial.ttf", 28); err != nil {
		panic(err)
	}

	dc.DrawStringAnchored("A", x, y, 0.5, 0.5)
	dc.SavePNG("out.png")
}

func random_position() (float64, float64) {
	x := radio + rand.Float64()*((screen_size-radio)-radio)
	y := radio + rand.Float64()*((screen_size-radio)-radio)
	return x, y
}

func check_distance(x1, y1, x2, y2 int) int {

	distanceX := x1 - x2
	distanceY := y1 - y2
	distance := distanceX*distanceX + distanceY*distanceY

	return distance
}*/
