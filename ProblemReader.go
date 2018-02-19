package main

import (
	"io/ioutil"
	"strings"
	"strconv"
)

func parseLine(line string) (int, float64, float64) {
	l := strings.Split(line, " ")
	id, _ := strconv.Atoi(l[0])
	x,  _ := strconv.ParseFloat(l[1], 64)
	y, _ := strconv.ParseFloat(l[1], 64)
	return id, x, y
}

func readProblemFile(path string) []City {
	data, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(data), "\n")
	citiesString := strings.Trim(strings.Split(lines[2], ": ")[1], "\r")
	numCities, _ := strconv.Atoi(citiesString)
	cityLines := lines[5: 5 + numCities]
	cities := make([]City, numCities)
	for i := 0; i < len(cityLines); i++ {
		id, x, y := parseLine(cityLines[i])
		cities[i] = City{id, x, y}
	}
	return cities
}