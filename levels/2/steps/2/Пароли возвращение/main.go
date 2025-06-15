package main

import (
	"unicode"
)

const minimumLength = 8

func ratePassword(pass string) string {
	hasLower := hasLowerCase(pass)
	hasUpper := hasUpper(pass)
	hasMinLength := hasMinimumLength(pass, minimumLength)

	switch {
	case hasLower && hasUpper && hasMinLength:
		return "Сложный пароль"
	case hasLower && hasUpper || hasLower && hasMinLength || hasUpper && hasMinLength:
		return "Средний пароль"
	case hasLower || hasUpper || hasMinLength:
		return "Слабый пароль"
	default:
		return "Пароль недостаточно безопасен, придумайте новый"
	}
}

func hasMinimumLength(text string, minLength int) bool {
	return len(text) >= minLength
}

func hasLowerCase(text string) bool {
	for _, c := range text {
		if unicode.IsLower(c) && unicode.IsLetter(c) {
			return true
		}
	}

	return false
}

func hasUpper(text string) bool {
	for _, c := range text {
		if unicode.IsUpper(c) && unicode.IsLetter(c) {
			return true
		}
	}

	return false
}
