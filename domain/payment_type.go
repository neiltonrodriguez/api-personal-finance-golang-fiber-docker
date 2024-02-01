package domain

type PaymentTypes []PaymentType

type PaymentType struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Descritpion string `json:"description"`
}
