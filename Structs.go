package main

type City struct {
	id int
	x, y float64
}

type Solution struct {
	cities [] City
	Distance float64
}

// For sorting by distance
type ByDistance []Solution

func (s ByDistance) Len() int {
	return len(s)
}

func (s ByDistance) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByDistance) Less(i, j int) bool {
	return s[i].Distance < s[j].Distance
}