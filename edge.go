package main

type Edge struct {
	name   string // unique id
	begin  Node
	end    Node
	weight float64
}


func (e Edge) GetName() string {
  return e.name
}

func (e Edge) GetWeight() float64 {
  return e.weight
}

// return source and destination nodes.
func (e Edge) GetNodes() (Node, Node) {
  return e.begin, e.end
}
