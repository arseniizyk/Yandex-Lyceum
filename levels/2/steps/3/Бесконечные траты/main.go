package main

import (
	"errors"
)

var Balance float64

var (
	ErrIncorrectBalance = errors.New("balance is incorrect")
	ErrIncorrectAmount  = errors.New("amount is incorrect")
)

func topUpBalance(amount float64) error {
	switch {
	case !isAmountValid(amount):
		return ErrIncorrectAmount
	case Balance < 0:
		return ErrIncorrectBalance
	}

	Balance += amount

	return nil
}

func chargeFromBalance(amount float64) error {
	switch {
	case !isAmountValid(amount):
		return ErrIncorrectAmount
	case Balance < amount:
		return ErrIncorrectBalance
	}

	Balance -= amount

	return nil
}

func TopUpAndGetBalance(amount float64) (float64, error) {
	err := topUpBalance(amount)
	if err != nil {
		return 0, err
	}
	return Balance, nil
}

func ChargeFromAndGetBalance(amount float64) (float64, error) {
	err := chargeFromBalance(amount)
	if err != nil {
		return 0, err
	}

	return Balance, nil
}

func isAmountValid(amount float64) bool {
	return amount > 0
}
