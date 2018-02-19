package main

import (
	"math"
)

func distanceBetweenPoints(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y - p2.y, 2))
}
