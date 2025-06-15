package main

func FiveSteps(array [5]int) [5]int {
	var result [5]int
	for i, j := 0, len(array)-1; i <= j; i, j = i+1, j-1 {
		result[j], result[i] = array[i], array[j]
	}
	return result
}
