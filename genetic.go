package main

type Individual struct {
	startNode Node
	endNode   Node
	path      []Node
	fitness   float64
}

type Population struct {
	individuals []Individual
}

func (g *Graph) geneticAlgorithm() ([]string, float64) {
	population := Population{}
	population.generatePopulation(100)

	return make([]string, 0), 0.0
}

func (p *Population) generatePopulation(nPopulation int) {

}

func (i *Individual) calculateFitness() {

}

func (i *Individual) makeCrossOver() {

}

func (i *Individual) makeMutation(second *Individual) {

}
