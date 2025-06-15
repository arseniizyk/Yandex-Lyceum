package main

import (
	"unicode"
)

const minimumLength = 8

func checkPassword(pass string) bool {
	if !hasMinimumLength(pass, minimumLength) || !hasUpper(pass) {
		return false
	}
	return true
}

func hasMinimumLength(text string, minLength int) bool {
	return len(text) >= minLength
}

func hasUpper(text string) bool {
	for _, c := range text {
		if unicode.IsUpper(c) && unicode.IsLetter(c) {
			return true
		}
	}

	return false
}
