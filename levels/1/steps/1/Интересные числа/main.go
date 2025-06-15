package main

import (
	"fmt"
	"math"
)

func main() {
	var number float64
	fmt.Scan(&number)

	sqrt := math.Sqrt(number)
	if math.IsNaN(sqrt) {
		fmt.Println(-1)
		return
	}
	fmt.Printf("%.3f\n", sqrt)
}
