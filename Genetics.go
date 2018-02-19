package main

import (
	"math/rand"
	"sort"
	"fmt"
)

// Mutate a solution
func (solution *Solution) mutate() {
	i, j := rand.Intn(len(solution.cities)), rand.Intn(len(solution.cities))
	switch rand.Intn(2) {
	case 0:		// Swap mutation
		solution.cities[i], solution.cities[j] = solution.cities[j], solution.cities[i]
	case 1:		// Reverse mutation
		indices := []int{i, j}
		sort.Ints(indices)
		for i := indices[1]/2-1; i >= indices[0]; i-- {
			opp := indices[1]-1-i
			solution.cities[i], solution.cities[opp] = solution.cities[opp], solution.cities[i]
		}
	}
	solution.Distance = calculateDistance(solution.cities)
}

func crossover(mama, papa *Solution)  {
	cities1 := make([]City, len(mama.cities))
	cities2 := make([]City, len(papa.cities))
	copy(cities1, mama.cities)
	copy(cities2, papa.cities)

}