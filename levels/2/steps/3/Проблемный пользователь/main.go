package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidInput    = errors.New("некорректный ввод")
	ErrInvalidAge      = errors.New("невалидный возраст")
	ErrInvalidName     = errors.New("невалидное имя")
	ErrInvalidPassport = errors.New("невалидная серия и номер паспорта")
)

func main() {
	var age int
	var name, passportSeriesAndNumber string

	_, err := fmt.Scanln(&age, &name, &passportSeriesAndNumber)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", ErrInvalidInput)
		return
	}

	if !isAgeValid(age) {
		fmt.Printf("Ошибка: %v\n", ErrInvalidAge)
		return
	}

	if !isNameValid(name) {
		fmt.Printf("Ошибка: %v\n", ErrInvalidName)
		return
	}

	if !isPassportSeriesAndNumberValid(passportSeriesAndNumber) {
		fmt.Printf("Ошибка: %v\n", ErrInvalidPassport)
		return
	}

	fmt.Printf("Имя: %s. Возраст: %d. Серия и номер паспорта: %s\n", name, age, passportSeriesAndNumber)
}

func isAgeValid(age int) bool {
	return age > 14 && age < 150
}

func isNameValid(name string) bool {
	return len(name) > 2
}

func isPassportSeriesAndNumberValid(seriesAndNumber string) bool {
	return len(seriesAndNumber) == 10
}
