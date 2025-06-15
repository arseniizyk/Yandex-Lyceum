package main

import (
	"fmt"
	"unicode/utf8"
)

func CountLengthAndBytes(first, second string) string {
	combined := fmt.Sprintf("%s%s", first, second)
	result := fmt.Sprintf("Объединённая строка: %s. Количество байт: %d. Количество символов: %d.", combined, len(combined), utf8.RuneCountInString(combined))

	return result
}
