package main

import (
	"fmt"
	"strings"
)

func main() {
	var number int
	fmt.Scan(&number)

	result := strings.Builder{}
	result.Write([]byte("Число "))

	if number == 0 {
		result.Write([]byte("0"))
	} else if number < 10 && number > -10 {
		result.Write([]byte("однозначное"))
	} else if number%2 == 0 {
		result.Write([]byte("чётное"))
	} else if number > 0 {
		result.Write([]byte("положительное"))
	} else if number < 0 {
		result.Write([]byte("красивое"))
	}

	fmt.Println(result.String())
}
