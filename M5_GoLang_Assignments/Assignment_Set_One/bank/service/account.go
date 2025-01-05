package service

import (
	"errors"
	"bank/model"
	"fmt"
)

var accounts []model.Account // Slice to store all accounts

func init() {
	// Pre-populate some accounts for testing
	accounts = []model.Account{
		{ID: 1, Name: "Alice", Balance: 500.0, TransactionHistory: []string{}},
		{ID: 2, Name: "Bob", Balance: 1000.0, TransactionHistory: []string{}},
	}
}

// DepositMoney adds money to an account
func DepositMoney(accountID int, amount float64) error {
	for i, acc := range accounts {
		if acc.ID == accountID {
			if amount <= 0 {
				return errors.New("deposit amount must be greater than zero")
			}
			accounts[i].Balance += amount
			accounts[i].TransactionHistory = append(accounts[i].TransactionHistory, fmt.Sprintf("Deposited $%.2f", amount))
			return nil
		}
	}
	return errors.New("account not found")
}

// WithdrawMoney withdraw1s money from an account
func WithdrawMoney(accountID int, amount float64) error {
	for i, acc := range accounts {
		if acc.ID == accountID {
			if amount <= 0 {
				return errors.New("withdrawal amount must be greater than zero")
			}
			if acc.Balance < amount {
				return errors.New("insufficient balance")
			}
			accounts[i].Balance -= amount
			accounts[i].TransactionHistory = append(accounts[i].TransactionHistory, fmt.Sprintf("Withdrew $%.2f", amount))
			return nil
		}
	}
	return errors.New("account not found")
}

// ViewBalance returns the balance of an account
func ViewBalance(accountID int) (float64, error) {
	for _, acc := range accounts {
		if acc.ID == accountID {
			return acc.Balance, nil
		}
	}
	return 0, errors.New("account not found")
}

// ViewTransactionHistory returns the transaction history of an account
func ViewTransactionHistory(accountID int) ([]string, error) {
	for _, acc := range accounts {
		if acc.ID == accountID {
			return acc.TransactionHistory, nil
		}
	}
	return nil, errors.New("account not found")
}
