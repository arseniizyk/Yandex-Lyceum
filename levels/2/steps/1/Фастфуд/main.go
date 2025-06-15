package main

import (
	"fmt"
)

func BuyFries(size string) {
	const smallPrice, mediumPrice, largePrice = 49, 89, 99

	switch size {
	case "S":
		printPrice(smallPrice, "Картошка фри")
	case "M":
		printPrice(mediumPrice, "Картошка фри")
	case "L":
		printPrice(largePrice, "Картошка фри")
	default:
		fmt.Println("Некорректный размер")
	}
}

func BuyCola(size string) {
	const smallPrice, mediumPrice, largePrice = 99, 119, 139

	switch size {
	case "S":
		printPrice(smallPrice, "Кола")
	case "M":
		printPrice(mediumPrice, "Кола")
	case "L":
		printPrice(largePrice, "Кола")
	default:
		fmt.Println("Некорректный размер")
	}
}

func printPrice(rubles int, product string) {
	fmt.Printf("%s будет стоить %d рублей\n", product, rubles)
}
