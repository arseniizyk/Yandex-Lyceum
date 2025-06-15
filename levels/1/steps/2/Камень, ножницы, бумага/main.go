package main

import "fmt"

func main() {
	var firstPlayerChoice, secondPlayerChoice string

	fmt.Scan(&firstPlayerChoice, &secondPlayerChoice)

	if firstPlayerChoice == "камень" && secondPlayerChoice == "камень" || firstPlayerChoice == "ножницы" && secondPlayerChoice == "ножницы" || firstPlayerChoice == "бумага" && secondPlayerChoice == "бумага" {
		fmt.Println("Ничья")
		return
	}

	if firstPlayerChoice == "камень" && secondPlayerChoice == "бумага" || firstPlayerChoice == "бумага" && secondPlayerChoice == "ножницы" || firstPlayerChoice == "ножницы" && secondPlayerChoice == "камень" {
		fmt.Println("Второй игрок победил")
		return
	}

	fmt.Println("Первый игрок победил")

}
