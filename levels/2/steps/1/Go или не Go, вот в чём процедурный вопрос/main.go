package main

import "fmt"

func GoOrNot(input string) {
	if input != "Go" {
		fmt.Println("Я знаю только Go!")
		return
	}
	fmt.Println("Go!")
}
