package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)

	input = strings.ReplaceAll(input, "а", "")
	input = strings.ReplaceAll(input, "о", "")

	fmt.Println(input)
}
