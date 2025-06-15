package main

import (
	"fmt"
)

func main() {
	var firstPass, secondPass string
	fmt.Scan(&firstPass, &secondPass)

	switch {
	case len(firstPass) >= 8 && len(secondPass) >= 8:
		fmt.Println("Оба пароля надёжные")
	case len(firstPass) >= 8:
		fmt.Println("Только первый пароль надёжный")
	case len(secondPass) >= 8:
		fmt.Println("Только второй пароль надёжный")
	default:
		fmt.Println("Оба пароля ненадёжные")
	}

}
