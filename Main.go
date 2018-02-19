package main

import (
	"time"
	"math/rand"
	"fmt"
)


func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	cities := readProblemFile("data/6.txt")
	population := make([]Solution, 5)
	for i := 0; i < len(population); i++ {
		population[i] = createRandomSolution(cities)
	}


	//for i := range population {
	//	population[i].mutate()
	//}
	//sort.Sort(ByDistance(population))

	s1 := createRandomSolution(cities)
	s2 := createRandomSolution(cities)

	fmt.Println("S1", s1)
	fmt.Println("S2", s2)

	crossover(&s1, &s2)


}