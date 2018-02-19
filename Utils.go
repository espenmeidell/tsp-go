package main

import (
	"math"
	"math/rand"
)

func distanceBetweenPoints(p1, p2 *City) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y - p2.y, 2))
}


func calculateDistance(cities []City) float64 {
	total := 0.0
	for i := 0; i < len(cities) - 1; i++ {
		total += distanceBetweenPoints(&cities[i], &cities[i+1])
	}
	total += distanceBetweenPoints(&cities[len(cities)-1], &cities[0])
	return total
}


func createRandomSolution(cities []City) Solution {
	order := make([]City, len(cities))
	for i, j := range rand.Perm(len(cities)){
		order[i] = cities[j]
	}
	return Solution{order, calculateDistance(order)}
}



//func populationStats(population *Population) (float64, float64, float64) {
//	best, worst, sum := population.solutions[0].distance, population.solutions[0].distance, 0.0
//	for _, p := range population.solutions{
//		if p.distance < best {
//			best = p.distance
//		}
//		if p.distance > worst {
//			worst = p.distance
//		}
//		sum += p.distance
//	}
//	average := sum / float64(len(population.solutions))
//	return best, worst, average
//}

func averageDist(solutions []Solution) float64 {
	sum := 0.0
	for _, s := range solutions {
		sum += s.Distance
	}
	return sum / float64(len(solutions))
}