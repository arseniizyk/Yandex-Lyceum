package main

import "fmt"

func main() {
	var firstNumber, secondNumber int
	fmt.Scan(&firstNumber, &secondNumber)
	if firstNumber > secondNumber {
		fmt.Println("Первое число больше второго")
	} else if firstNumber < secondNumber {
		fmt.Println("Второе число больше первого")
	} else {
		fmt.Println("Числа равны")
	}
}
