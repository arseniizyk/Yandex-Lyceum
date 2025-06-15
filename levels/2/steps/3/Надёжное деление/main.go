package main

import (
	"errors"
)

var (
	DivisionByZeroError = "division by zero is not allowed"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New(DivisionByZeroError)
	}
	return a / b, nil
}
