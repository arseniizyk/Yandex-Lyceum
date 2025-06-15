package main

func SumOfValuesInMap(m map[int]int) int {
	total := 0

	for _, v := range m {
		total += v
	}

	return total
}
