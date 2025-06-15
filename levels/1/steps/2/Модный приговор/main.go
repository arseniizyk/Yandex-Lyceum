package main

import "fmt"

func main() {
	var degrees int
	var abs string
	fmt.Scan(&abs)
	if abs == "0" {
		fmt.Println("Стоит надеть куртку")
		return
	}
	fmt.Scan(&degrees)

	switch abs {
	case "+":
		if degrees > 20 {
			fmt.Println("Стоит надеть майку и шорты")
			return
		} else if degrees >= 10 && degrees <= 20 {
			fmt.Println("Стоит надеть штаны и кофту")
			return
		}
		fmt.Println("Стоит надеть куртку")

	case "-":
		if degrees < 5 {
			fmt.Println("Стоит надеть куртку")
			return
		}
		fmt.Println("Стоит надеть зимнюю куртку")
	}

}
