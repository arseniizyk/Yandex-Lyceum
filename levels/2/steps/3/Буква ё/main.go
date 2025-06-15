package main

import (
	"fmt"
	"strings"
)

func CheckLetters(text string) string {
	number := strings.Count(text, "е")
	if number != 0 {
		return fmt.Sprintf("Количество возможных ошибок: %d, перепроверьте текст", number)
	}
	return "Текст готов к публикации!"
}
