package main

import (
	"fmt"
	"math"
)

func main() {
	var heightFirst, heightSecond, heightThird float64
	fmt.Scan(&heightFirst, &heightSecond, &heightThird)
	fmt.Println(math.Min(math.Min(heightFirst, heightSecond), heightThird))
}
