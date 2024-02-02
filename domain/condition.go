package domain

type Conditions []Condition

type Condition struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Descritpion string `json:"description"`
}
