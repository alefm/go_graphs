package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type GeneticFrontendState struct {
	SolutionPath    []string	`json:"name,omitempty"`
	Fitness			float64 	`json:"name,omitempty"`
	Permutation		int			`json:"name,omitempty"`
	Population		int			`json:"name,omitempty"`
	Selected		int			`json:"name,omitempty"`
	NodePath		[]string	`json:"name,omitempty"`
}


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
		if individual.fitness < bestIndividual.fitness {
			bestIndividual = individual
		}
	}

	return bestIndividual
}

func (g *Graph) runGeneticAlgorithm(population Population, stopCriterion int, coRatio float64, mutationRatio float64, nPopulation int) Individual {
	nCrossOver := int((coRatio / 100) * float64(nPopulation))
	nMutations := int((mutationRatio / 100) * float64(nPopulation))

	population.makeTournament(nPopulation)

	// Generate fitness of all individuals
	for key, ind := range population.individuals {
		g.calculateFitness(&ind)
		population.individuals[key].fitness = ind.fitness
	}

	// population.makeTournament(nPopulation)
	if stopCriterion == 0 {
		return population.getBestIndividual()
	}

	// Generate CrossOver Population
	for i := 0; i < nCrossOver; i++ {
		indexA := rand.Intn(len(population.individuals))
		indexB := rand.Intn(len(population.individuals))
		population.individuals[indexA].makeCrossOver(&population, &population.individuals[indexB])
	}

	// Generate Mutation Population
	for i := 0; i < nMutations; i++ {
		index := rand.Intn(len(population.individuals))
		population.individuals[index].makeMutation()
	}

	return g.runGeneticAlgorithm(population, stopCriterion-1, coRatio, mutationRatio, nPopulation)
}

func (g *Graph) geneticAlgorithm(startNode string, stopCriterion int, coRatio float64, mutationRatio float64, nPopulation int) ([]string, float64) {
	population := Population{}
	population.startNode = *g.GetNode(startNode)
	population.individuals = make([]Individual, 0)
	rand.Seed(time.Now().Unix())

	nodeList := g.NodeListAsString()

	// remove the startNode from the list
	nodeList = findAndRemove(nodeList, startNode)
	g.generatePopulation(&population, nodeList, nPopulation)

	solution := g.runGeneticAlgorithm(population, stopCriterion, coRatio, mutationRatio, nPopulation)
	fmt.Println("Solução encontrada", solution.path, "com o fitness", solution.fitness)
	g.frontend.genetic.SolutionPath = solution.path
	g.frontend.genetic.Fitness = solution.fitness

	return solution.path, solution.fitness
}

/*func (g *Graph) calculateFitness(i *Individual) {
	distance := 0.0
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
		distance = distance + g.calculateDistancePythagorean(source, dest)
	}

	i.fitness = 1 / distance
}*/

func (g *Graph) calculateFitness(i *Individual) {
	distance := 0.0
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

		if g.isAdjacent(source, dest) {
			distance += g.EdgeList[g.GetEdgeIndex(source.GetName(), dest.GetName())].weight
		} else {
			distance += g.getMaxWeight() * 5
		}

	}

	i.fitness = distance
}

func (p *Population) makeTournament(nPopulation int) {
	k := 0.75
	newIndividuals := make([]Individual, 0)
	for i := 0; i < nPopulation; i++ {
		indexA := rand.Intn(len(p.individuals))
		indexB := rand.Intn(len(p.individuals))
		var best int
		var worse int
		r := rand.Float64()

		if p.individuals[indexA].fitness > p.individuals[indexB].fitness {
			best = indexA
			worse = indexB
		} else {
			best = indexB
			worse = indexA
		}

		if r < k {
			newIndividuals = append(newIndividuals, p.individuals[best])
		} else {
			newIndividuals = append(newIndividuals, p.individuals[worse])
		}
	}

	p.individuals = newIndividuals
}

func (g *Graph) generatePopulation(p *Population, nodeList []string, nPopulation int) {
	fmt.Println("Gerando população atraves dos nodos: ", nodeList)
	g.frontend.genetic.NodePath = nodeList

	permutation := make([]Individual, 0)
	Perm(nodeList, func(a []string) {
		// append the starting node at the first position
		a = append([]string{p.startNode.Name}, a...)
		individual := Individual{a, 0}
		g.calculateFitness(&individual)
		permutation = append(permutation, individual)
	})
	p.individuals = permutation



	fmt.Println("Quantidade de permutacoes geradas ", len(p.individuals))
	g.frontend.genetic.Permutation = len(p.individuals)

	fmt.Println("Escolhendo", nPopulation, " delas...")
	g.frontend.genetic.Selected = nPopulation

	p.makeTournament(nPopulation)

	fmt.Println("População inicial gerada! Com o tamanho: ", len(p.individuals))
	g.frontend.genetic.Population = len(p.individuals)
}

func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func existInSlice(slice []string, key string) bool {

	for _, sliceKey := range slice {
		if sliceKey == key {
			return true
		}
	}

	return false
}
func fillSlice(slice []string, parent []string) []string {
	for _, key := range parent {
		if !existInSlice(slice, key) {
			slice = append(slice, key)
		}
	}

	return slice
}

func (i *Individual) makeCrossOver(population *Population, second *Individual) {
	size := len(i.path)
	halfSize := int(size / 2)
	remainSize := (size - halfSize) / 2

	childA := make([]string, 0)
	childB := make([]string, 0)

	childA = append(childA, i.path[:remainSize]...)
	childA = append(childA, second.path[remainSize:size-remainSize]...)
	childA = append(childA, i.path[size-remainSize:size]...)

	childB = append(childB, second.path[:remainSize]...)
	childB = append(childB, i.path[remainSize:size-remainSize]...)
	childB = append(childB, second.path[size-remainSize:size]...)

	childA = unique(childA)
	childB = unique(childB)

	childA = fillSlice(childA, i.path)
	childB = fillSlice(childB, second.path)

	population.individuals = append(population.individuals, Individual{childA, 0})
	population.individuals = append(population.individuals, Individual{childB, 0})
}

func (i *Individual) makeMutation() {
	indexA := rand.Intn(len(i.path))
	indexB := rand.Intn(len(i.path))

	if indexA == 0 {
		indexA = 1
	}

	if indexB == 0 {
		indexB = 1
	}

	aux := i.path[indexB]
	i.path[indexB] = i.path[indexA]
	i.path[indexA] = aux
}
