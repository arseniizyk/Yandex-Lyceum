package main

import (
	"fmt"
	"time"
)

func main() {
	var input string

	fmt.Scan(&input)

	layout := "2006-01-02/15:04:05"

	timeInput, err := time.Parse(layout, input)
	if err != nil {
		panic(err)
	}

	output := fmt.Sprintf("Текущее время %d часов, %d минут. Ты точно не забыл про важные дела на сегодня?", timeInput.Hour(), timeInput.Minute())

	fmt.Println(output)
}
