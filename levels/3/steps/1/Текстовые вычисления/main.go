package main

import (
	"strings"
)

var replacement = map[rune]string{
	'0': "ноль",
	'1': "один",
	'2': "два",
	'3': "три",
	'4': "четыре",
	'5': "пять",
	'6': "шесть",
	'7': "семь",
	'8': "восемь",
	'9': "девять",
	'+': "плюс",
	'-': "минус",
	'*': "умножить на",
	'/': "разделить на",
}

func NumbersToLetters(input string) string {
	var result strings.Builder

	for _, c := range input {
		if val, ok := replacement[c]; ok {
			result.WriteString(val)
			continue
		} else {
			result.WriteRune(c)
		}
	}

	return result.String()
}
