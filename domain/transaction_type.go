package domain

type TransactionTypes []TransactionType

type TransactionType struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Descritpion     string    `json:"description"`
}