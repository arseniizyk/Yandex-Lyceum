package main

import (
	"math"
)

func FindMinMaxInArray(array [10]int) (int, int) {
	max := math.MinInt
	min := math.MaxInt
	for _, num := range array {
		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
	}

	return max, min
}
