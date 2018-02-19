package main

import (
	"math/rand"
)

// Mutate a solution
func (solution *Solution) mutate() {
	//fmt.Println("original:", s.distance)
	switch rand.Intn(1) {
	case 0:		// Swap mutation
		i, j := rand.Intn(len(solution.cities)), rand.Intn(len(solution.cities))
		solution.cities[i], solution.cities[j] = solution.cities[j], solution.cities[i]
	}
	solution.Distance = calculateDistance(solution.cities)
	//fmt.Println("new:", s.distance)

}