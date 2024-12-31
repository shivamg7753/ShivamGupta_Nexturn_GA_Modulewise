package main

import (
	"bank/service"
	"fmt"
)

func main() {
	for {
		fmt.Println("Bank Transaction System")
		fmt.Println("1. Deposit Money")
		fmt.Println("2. Withdraw Money")
		fmt.Println("3. View Balance")
		fmt.Println("4. View Transaction History")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1: // Deposit Money
			var accountID int
			var amount float64
			fmt.Print("Enter account ID: ")
			fmt.Scan(&accountID)
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&amount)
			err := service.DepositMoney(accountID, amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful.")
			}
		case 2: // Withdraw Money
			var accountID int
			var amount float64
			fmt.Print("Enter account ID: ")
			fmt.Scan(&accountID)
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scan(&amount)
			err := service.WithdrawMoney(accountID, amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful.")
			}
		case 3: // View Balance
			var accountID int
			fmt.Print("Enter account ID: ")
			fmt.Scan(&accountID)
			balance, err := service.ViewBalance(accountID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Balance: $%.2f\n", balance)
			}
		case 4: // View Transaction History
			var accountID int
			fmt.Print("Enter account ID: ")
			fmt.Scan(&accountID)
			history, err := service.ViewTransactionHistory(accountID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Transaction History:")
				for _, entry := range history {
					fmt.Println(entry)
				}
			}
		case 5: // Exit
			fmt.Println("Exiting system. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
