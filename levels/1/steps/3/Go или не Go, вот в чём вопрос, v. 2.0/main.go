package main

import "fmt"

func main() {
	for range 5 {
		var input string
		fmt.Scanln(&input)
		if input == "Go" {
			fmt.Println("Go!")
		} else {
			fmt.Println("Я знаю только Go!")
		}
	}
}
