package main

import "fmt"

func main() {
	cities := readProblemFile("data/1.txt")
	s1 := createRandomSolution(cities)
	s2 := createRandomSolution(cities)

	fmt.Println(calculateDistance(&s1))
	fmt.Println(calculateDistance(&s2))
}