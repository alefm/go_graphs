package main

/* verify if a given Node have any connection with another node in list */
func (g *Graph) connectionInSlice(node Node, list []Node) bool {
	connection := false

	for _, nodeInList := range list {
		if nodeInList.GetName() == node.GetName() || g.isAdjacent(nodeInList, node) {
			connection = true
			break
		}
	}

	return connection
}

// Coloring should put colors in all nodes in graph
func (g *Graph) Coloring() {
	pallete := [...]string{"gold", "green", "hotpink", "orchid", "red", "blue", "tan", "yellow", "magenta", "cyan", "blueviolet", "olivedrab3"}
	colorMap := make(map[int][]Node)
	colorIdx := 0

	var slice []Node
	colorMap[colorIdx] = slice

	for idx, nodeA := range g.NodeList {

		colored := false
		for i := 0; i <= colorIdx; i++ {
			if value, ok := colorMap[i]; ok {

				if !g.connectionInSlice(nodeA, value) {
					nodeA.SetColor(pallete[i])
					g.NodeList[idx] = nodeA
					value = append(value, nodeA)
					colorMap[i] = value
					colored = true
					break
				}

			}
		}

		if !colored {
			colorIdx = colorIdx + 1
			nodeA.SetColor(pallete[colorIdx])
			g.NodeList[idx] = nodeA
			var newSlice []Node
			newSlice = append(newSlice, nodeA)
			colorMap[colorIdx] = newSlice
		}
	}
}
