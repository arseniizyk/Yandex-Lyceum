package main

import "fmt"

func main() {
	var products int
	var discount, total float64
	fmt.Scan(&products, &discount)
	discount /= 100
	for range products {
		var price float64
		fmt.Scanln(&price)
		if discount != 0 {
			total += price * (1 - discount)
			continue
		}
		total += price
	}
	fmt.Println(total)
}
