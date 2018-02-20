package main

import (
	"time"
	"math/rand"
	"fmt"
)


func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	cities := readProblemFile("data/1.txt")
	population := make([]Solution, 10)
	for i := 0; i < len(population); i++ {
		population[i] = createRandomSolution(cities)
	}


	for i := 0; i < 100000; i++ {
		population = evolve(population)
		fmt.Println(population[0].Distance)
	}


}