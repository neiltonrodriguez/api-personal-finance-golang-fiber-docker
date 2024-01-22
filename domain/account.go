package domain

import "time"

type AccountInput struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Balance     float64   `json:"balance"`
	TypeAccount int       `json:"type_account"`
	BankId      int       `json:"bank_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AccountOutput struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Balance     float64   `json:"balance"`
	TypeAccount string    `json:"type_account"`
	Bank        string    `json:"bank"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BalanceTotal struct {
	Balance float64 `json:"balance"`
}
