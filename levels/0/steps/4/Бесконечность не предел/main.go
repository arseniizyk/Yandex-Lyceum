package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Scan(&a, &b)

	max := math.Max(a, b)
	fmt.Println(math.Round(max))
}
