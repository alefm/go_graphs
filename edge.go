package main

type Edge struct {
	// Unique ID
	name   string
	src    Node
	dst    Node
	weight float64
}

