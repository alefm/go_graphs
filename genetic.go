package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Individual struct {
	path    []string
	fitness float64
}

type Population struct {
	startNode   Node
	individuals []Individual
}

func (g *Graph) calculateDistancePythagorean(source Node, destination Node) float64 {
	x := math.Pow((destination.GraphPoint.X - source.GraphPoint.Y), 2)
	y := math.Pow((destination.GraphPoint.Y - source.GraphPoint.Y), 2)
	return math.Sqrt((x + y))
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

func (p *Population) getBestIndividual() Individual {
	bestIndividual := p.individuals[0]

	for _, individual := range p.individuals {
		if individual.fitness > bestIndividual.fitness {
			bestIndividual = individual
		}
	}

	return bestIndividual
}

func (g *Graph) runGeneticAlgorithm(population Population, stopCriterion int) Individual {
	// Generate fitness of all individuals
	for _, ind := range population.individuals {
		g.calculateFitness(&ind)
	}

	if stopCriterion == 0 {
		return population.getBestIndividual()
	}

	// sort crossover and make it

	// sort mutation and make it

	return g.runGeneticAlgorithm(population, stopCriterion-1)
}

func (g *Graph) geneticAlgorithm(startNode string, stopCriterion int, coRatio float64, mutationRatio float64) ([]string, float64) {
	population := Population{}
	population.startNode = *g.GetNode(startNode)
	population.individuals = make([]Individual, 0)
	rand.Seed(time.Now().Unix())

	nodeList := g.NodeListAsString()

	// remove the startNode from the list
	nodeList = findAndRemove(nodeList, startNode)
	population.generatePopulation(nodeList, 100)

	solution := g.runGeneticAlgorithm(population, stopCriterion)
	fmt.Println("Solução encontrada", solution.path, "com o fitness", solution.fitness)
	return solution.path, solution.fitness
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

	//TODO: alterar para torneio
	for i := 0; i < nPopulation; i++ {
		index := rand.Intn(len(permutation))
		individual := Individual{permutation[index], 0}
		p.individuals = append(p.individuals, individual)
	}

	fmt.Println("População inicial gerada! Com o tamanho: ", len(p.individuals))
}

func (g *Graph) calculateFitness(i *Individual) {
	for index := range i.path {
		var source Node
		var dest Node
		if index+1 > len(i.path)-1 {
			source = *g.GetNode(i.path[index])
			dest = *g.GetNode(i.path[0])
		} else {
			source = *g.GetNode(i.path[index])
			dest = *g.GetNode(i.path[index+1])
		}
		i.fitness = 1 / g.calculateDistancePythagorean(source, dest)
	}
	fmt.Println(i.path, " : ", i.fitness)
}

func (i *Individual) makeCrossOver(second *Individual) {

}

func (i *Individual) makeMutation() {
	indexA := rand.Intn(len(i.path))
	indexB := rand.Intn(len(i.path))

	aux := i.path[indexB]
	i.path[indexB] = i.path[indexA]
	i.path[indexA] = aux
}
