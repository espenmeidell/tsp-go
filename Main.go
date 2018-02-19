package main

import "fmt"

func main() {
	cities := readProblemFile("data/1.txt")
	test := Solution{cities}
	fmt.Println(calculateDistance(&test))
}