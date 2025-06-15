package main

import "fmt"

func main() {
	var students int
	fmt.Scanln(&students)
	for range students {
		var score float64
		fmt.Scanln(&score)
		if score >= 90 && score <= 100 {
			fmt.Println(5)
		} else if score >= 75 && score <= 89 {
			fmt.Println(4)
		} else if score >= 50 && score <= 74 {
			fmt.Println(3)
		} else if score >= 0 && score <= 49 {
			fmt.Println(2)
		} else {
			fmt.Println("Неверный балл")
		}
	}
}
