package main

import "math"

func SquareRoots(a, b, c float64) (float64, float64) {
	disc := findDiscriminant(a, b, c)
	if disc < 0 {
		return 0, 0
	}
	return (-b - math.Sqrt(disc)) / (2 * a), (-b + math.Sqrt(disc)) / (2 * a)
}

func findDiscriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}
