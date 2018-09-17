package main

type Node struct {
	// Unique ID
	name string
}

// NewNode returns a new Node.
func NewNode(id string) *Node {
	return &Node{
		name: id,
	}
}
