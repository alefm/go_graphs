package main

type Edge struct {
	// Unique ID
	name   string
	src    Node
	dst    Node
	weight float64
}

// NewEdge returns a new Edge.
func NewEdge(id string, src, dst Node, weight float64) *Edge {
	return &Edge{
		name:   id,
		src:    src,
		dst:    dst,
		weight: weight,
	}
}
