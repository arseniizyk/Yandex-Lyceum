package main

import "fmt"

type Account struct {
	balance float64
	owner   string
}

func NewAccount(owner string) *Account {
	return &Account{
		owner:   owner,
		balance: 0,
	}
}

var (
	ErrNegativeValue    = fmt.Errorf("value can not be less than 0")
	ErrUnsufficentFunds = fmt.Errorf("unsufficient funds to complete this operation")
)

func (a *Account) SetBalance(value float64) error {
	if value <= 0 {
		return ErrNegativeValue
	}

	a.balance = value
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(value float64) error {
	if value <= 0 {
		return ErrNegativeValue
	}

	a.balance += value
	return nil
}

func (a *Account) Withdraw(value float64) error {
	if value <= 0 {
		return ErrNegativeValue
	}
	if value > a.balance {
		return ErrUnsufficentFunds
	}

	a.balance -= value
	return nil
}
