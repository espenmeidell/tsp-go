package main

import (
	"math/rand"
	"sort"
)

// Mutate a solution
func (solution *Solution) mutate() {
	i, j := rand.Intn(len(solution.cities)), rand.Intn(len(solution.cities))
	switch rand.Intn(2) {
	case 0: // Swap mutation
		solution.cities[i], solution.cities[j] = solution.cities[j], solution.cities[i]
	case 1: // Reverse mutation
		indices := []int{i, j}
		sort.Ints(indices)
		for i := indices[1]/2 - 1; i >= indices[0]; i-- {
			opp := indices[1] - 1 - i
			solution.cities[i], solution.cities[opp] = solution.cities[opp], solution.cities[i]
		}
	}
	solution.Distance = calculateDistance(solution.cities)
}

func crossover(mama, papa *Solution) (Solution, Solution) {
	cities1 := make([]City, len(mama.cities))
	cities2 := make([]City, len(papa.cities))
	copy(cities1, mama.cities)
	copy(cities2, papa.cities)
	child1 := Solution{cities1, calculateDistance(cities1)}
	child2 := Solution{cities2, calculateDistance(cities2)}
	child1.mutate()
	child2.mutate()
	return child1, child1
}

func getParents(population []Solution) (Solution, Solution) {
	tournament := []Solution{
		population[rand.Intn(len(population))],
		population[rand.Intn(len(population))],
		population[rand.Intn(len(population))],
		population[rand.Intn(len(population))]}
	sort.Sort(ByDistance(tournament))
	return tournament[0], tournament[1]
}

func evolve(population []Solution) []Solution {
	newPopulation := make([]Solution, 0)
	for len(newPopulation) < len(population) {
		mama, papa := getParents(population)
		child1, child2 := crossover(&mama, &papa)
		child1.mutate()
		child2.mutate()
		newPopulation = append(newPopulation, child1)
		newPopulation = append(newPopulation, child2)
	}
	newPopulation = append(newPopulation, population...)
	sort.Sort(ByDistance(newPopulation))
	return newPopulation[:len(population)]
}
