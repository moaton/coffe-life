package domain

import "github.com/google/uuid"

type Order struct {
	ID          uuid.UUID `json:"id"`
	Fullname    string    `json:"fullname"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
}

type Orders []Order

type OrdersFoods struct {
	OrderID uuid.UUID
	FoodID  uuid.UUID
}
