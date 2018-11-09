package main

// Node Structure
type Node struct {
	// Unique ID
	Name       string `json:"name,omitempty"`
	Color      string `json:"color,omitempty"`
	GraphPoint Point  `json:"graph_point,omitempty"`
}

// GetName - return current node name
func (n Node) GetName() string {
	return n.Name
}

// SetColor - set current node color
func (n *Node) SetColor(color string) {
	n.Color = color
}

// GetColor - get current node color
func (n Node) GetColor() string {
	if n.Color != "" {
		return n.Color
	}

	return "black"
}
