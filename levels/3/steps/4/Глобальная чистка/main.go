package main

func DeleteLongKeys(m map[string]int) map[string]int {
	cleaned := make(map[string]int)

	for k, v := range m {
		if len(k) < 6 {
			continue
		}

		cleaned[k] = v
	}

	return cleaned
}
