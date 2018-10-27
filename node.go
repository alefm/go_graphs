package main

// Node Structure
type Node struct {
	// Unique ID
	name string
	color string
}

// GetName - return current node name
func (n Node) GetName() string {
	return n.name
}

// SetColor - set current node color
func (n *Node) SetColor(color string){
	n.color = color
}

// GetColor - get current node color
func (n Node) GetColor() string {
	if n.color != "" {
		return n.color
	}

	return "black"
}
