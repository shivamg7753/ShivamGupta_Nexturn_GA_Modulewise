package model

type Account struct {
	ID                int      `json:"id"`
	Name              string   `json:"name"`
	Balance           float64  `json:"balance"`
	TransactionHistory []string `json:"transaction_history"`
}
