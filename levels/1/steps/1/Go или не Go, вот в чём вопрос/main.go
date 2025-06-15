package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	if input != "Go" {
		fmt.Println("Я знаю только Go!")
		return
	}
	fmt.Println("Go!")
}
