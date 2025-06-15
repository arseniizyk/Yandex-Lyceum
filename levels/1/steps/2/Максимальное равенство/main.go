package main

import "fmt"

func main() {
	var firstNumber, secondNumber, thirdNumber float64

	fmt.Scan(&firstNumber, &secondNumber, &thirdNumber)

	if firstNumber == secondNumber && secondNumber == thirdNumber {
		fmt.Println("Максимальное равенство")
		return
	}

	fmt.Println("Не равны")
}
