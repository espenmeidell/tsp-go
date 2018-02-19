package main

import (
	"math"
	"math/rand"
)

func distanceBetweenPoints(p1, p2 *City) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y - p2.y, 2))
}


func calculateDistance(solution *Solution) float64 {
	total := 0.0
	for i := 0; i < len(solution.cities) - 1; i++ {
		total += distanceBetweenPoints(&solution.cities[i], &solution.cities[i+1])
	}
	total += distanceBetweenPoints(&solution.cities[len(solution.cities)-1], &solution.cities[0])
	return total
}


func createRandomSolution(cities []City) Solution {
	order := make([]City, len(cities))
	for i, j := range rand.Perm(len(cities)){
		order[i] = cities[j]
	}
	return Solution{order}
}