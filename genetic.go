package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Individual struct {
	path    []string
	fitness float64
}

type Population struct {
	startNode   Node
	endNode     Node
	individuals []Individual
}

// Perm calls f with each permutation of a.
func Perm(a []string, f func([]string)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []string, f func([]string), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func findAndRemove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func (g *Graph) runGeneticAlgorithm(population Population) {
	// calculate fitness of all individuals

	// satisfy stop criterion?
	// then end the algorithm

	// sort crossover and make it

	// sort mutation and make it

	// runGeneticAlgorithm again
}

func (g *Graph) geneticAlgorithm(startNode string, endNode string) ([]string, float64) {
	population := Population{}
	population.startNode = *g.GetNode(startNode)
	population.endNode = *g.GetNode(endNode)
	population.individuals = make([]Individual, 0)

	nodeList := g.NodeListAsString()

	// remove the startNode from the list
	nodeList = findAndRemove(nodeList, startNode)
	population.generatePopulation(nodeList, 100)

	g.runGeneticAlgorithm(population)

	return make([]string, 0), 0.0
}

func (p *Population) generatePopulation(nodeList []string, nPopulation int) {
	fmt.Println("Gerando população atraves dos nodos: ", nodeList)

	permutation := make([][]string, 0)
	Perm(nodeList, func(a []string) {
		// append the starting node at the first position
		a = append([]string{p.startNode.Name}, a...)
		permutation = append(permutation, a)
	})
	fmt.Println("Quantidade de permutacoes geradas ", len(permutation))
	fmt.Println("Escolhendo", nPopulation, " delas...")
	rand.Seed(time.Now().Unix())

	for i := 0; i < nPopulation; i++ {
		index := rand.Intn(len(permutation))
		individual := Individual{permutation[index], 0}
		p.individuals = append(p.individuals, individual)
	}

	fmt.Println("População inicial gerada!")
}

func (i *Individual) calculateFitness() {

}

func (i *Individual) makeCrossOver() {

}

func (i *Individual) makeMutation(second *Individual) {

}
