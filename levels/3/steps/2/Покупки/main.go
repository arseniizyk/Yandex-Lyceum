package main

func SumOfArray(array [6]int) int {
	var total int
	for _, num := range array {
		total += num
	}

	return total
}
