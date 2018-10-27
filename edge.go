package main

// Edge structure
type Edge struct {
	name   string // unique id
	begin  Node
	end    Node
	weight float64
	color string
}

// GetName - return current edge name
func (e Edge) GetName() string {
	return e.name
}

// GetWeight - return current edge weight
func (e Edge) GetWeight() float64 {
	return e.weight
}

// GetNodes - return source and destination nodes.
func (e Edge) GetNodes() (Node, Node) {
	return e.begin, e.end
}

// SetColor - set current edge color
func (e *Edge) SetColor(color string){
	e.color = color
}

// GetColor - get current edge color
func (e Edge) GetColor() string {
	if e.color != "" {
		return e.color
	}

	return "black"
}
