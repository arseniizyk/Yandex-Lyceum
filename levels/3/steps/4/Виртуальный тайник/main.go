package main

import "math"

func FindMaxKey(m map[int]int) int {
	max := math.MinInt

	for k, _ := range m {
		if k > max {
			max = k
		}
	}

	return max
}
