package model

import "time"

type Guest struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	IdCardNumber string    `json:"id_card_number"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	CreatedAt    time.Time `json:"created_at"`
}
