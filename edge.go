package main

// Edge structure
type Edge struct {
	name   string // unique id
	begin  Node
	end    Node
	weight float64
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
