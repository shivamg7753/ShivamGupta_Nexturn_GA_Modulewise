package handler

import (
	// "bank/model"
	"bank/service"
	"fmt"
)

// DepositMoney handles deposit transactions
func DepositMoney() {
	var id int
	var amount float64

	fmt.Print("Enter account ID: ")
	fmt.Scan(&id)

	account, err := service.ValidateAccountID(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print("Enter amount to deposit: ")
	fmt.Scan(&amount)

	err = service.Deposit(account, amount)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Deposit successful!")
}

// WithdrawMoney handles withdrawal transactions
func WithdrawMoney() {
	var id int
	var amount float64

	fmt.Print("Enter account ID: ")
	fmt.Scan(&id)

	account, err := service.ValidateAccountID(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print("Enter amount to withdraw: ")
	fmt.Scan(&amount)

	err = service.Withdraw(account, amount)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Withdrawal successful!")
}

// ViewBalance displays account balance
func ViewBalance() {
	var id int

	fmt.Print("Enter account ID: ")
	fmt.Scan(&id)

	account, err := service.ValidateAccountID(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Account Balance: $%.2f\n", account.Balance)
}

// ViewTransactionHistory displays account transaction history
func ViewTransactionHistory() {
	var id int

	fmt.Print("Enter account ID: ")
	fmt.Scan(&id)

	account, err := service.ValidateAccountID(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Transaction History:")
	for _, transaction := range account.TransactionHist {
		fmt.Println(transaction)
	}
}
