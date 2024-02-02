package domain

import "time"

type Transaction struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Value     float64   `json:"float"`
	Category  string    `json:"category"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TransactionInput struct {
	Id                int       `json:"id"`
	Title             string    `json:"title"`
	Value             float64   `json:"value"`
	AccountId         int       `json:"account_id"`
	CategoryId        int       `json:"category_id"`
	TypeTransactionId int       `json:"type_transaction_id"`
	ConditionId       int       `json:"condition_id"`
	TypePaymentId     int       `json:"type_payment_id"`
	Descritpion       string    `json:"description"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type TransactionTotal struct {
	In  float64 `json:"total_in,omitempty"`
	Out float64 `json:"total_out,omitempty"`
}

type TransactionOutput struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Value           float64   `json:"value"`
	Category        string    `json:"category"`
	Condition       string    `json:"condition"`
	TypeTransaction string    `json:"type_transaction"`
	TypePayment     string    `json:"type_payment"`
	Descritpion     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
