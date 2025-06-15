package main

func SwapKeysAndValues(m map[string]string) map[string]string {
	swapped := make(map[string]string)

	for k, v := range m {
		swapped[v] = k
	}

	return swapped
}
