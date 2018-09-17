package main

type Node struct {
	// Unique ID
	name string
}

func (n Node) GetName() string {
  return n.name
}
