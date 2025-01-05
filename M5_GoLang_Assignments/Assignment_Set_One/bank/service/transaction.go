package service

import (
	"bank/model"
	"errors"
	"fmt"
)

// Deposit adds money to the account
func Deposit(account *model.Account, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	account.Balance += amount
	transaction := fmt.Sprintf("Deposited: $%.2f", amount)
	account.TransactionHistory = append(account.TransactionHistory, transaction)
	return nil
}

// Withdraw subtracts money from the account
func Withdraw(account *model.Account, amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}
	if account.Balance < amount {
		return errors.New("insufficient balance")
	}
	account.Balance -= amount
	transaction := fmt.Sprintf("Withdrew: $%.2f", amount)
	account.TransactionHistory = append(account.TransactionHistory, transaction)
	return nil
}
