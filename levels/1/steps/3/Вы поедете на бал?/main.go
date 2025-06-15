package main

import "fmt"

func main() {
	for {
		var answer string

		fmt.Scanln(&answer)

		switch answer {
		case "да", "чёрный", "белый", "нет":
			fmt.Println("Поражение")
			return
		default:
			fmt.Println("Игра продолжается")
		}
	}
}
