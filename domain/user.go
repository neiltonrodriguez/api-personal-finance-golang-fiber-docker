package domain

import "time"

type User struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Phone        string    `json:"phone"`
	Address      string    `json:"address"`
	Neighborhood string    `json:"neighborhood"`
	Zipcode      string    `json:"zipcode"`
	City         string    `json:"city"`
	UF           string    `json:"uf"`
	Number       string    `json:"number"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
