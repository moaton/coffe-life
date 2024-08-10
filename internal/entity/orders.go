package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID `json:"id" gorm:"default:uuid_generate_v4()"`
	Fullname    string    `json:"fullname"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type Orders []Order

type OrdersFoods struct {
	OrderID uuid.UUID
	FoodID  uuid.UUID
}
