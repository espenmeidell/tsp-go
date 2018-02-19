package main

import (
	//"sort"
	//"fmt"
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

	fmt.Println(averageDist(population))
	for i, _ := range population {
		population[i].mutate()
	}
	fmt.Println(averageDist(population))



}