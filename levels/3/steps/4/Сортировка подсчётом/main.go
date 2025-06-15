package main

func CountingSort(contacts []string) map[string]int {
	counted := make(map[string]int)

	for _, c := range contacts {
		counted[c]++
	}

	return counted
}
