package main

// Node Structure
type Node struct {
	// Unique ID
	name string
}

// GetName - return current node name
func (n Node) GetName() string {
	return n.name
}
