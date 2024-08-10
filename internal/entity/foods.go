package entity

import (
	"time"

	"github.com/google/uuid"
)

type Food struct {
	ID          uuid.UUID `json:"id" gorm:"default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Category    string    `json:"category"`
	Price       string    `json:"price"`
	IsNew       bool      `json:"is_new"`
	IsSpicy     bool      `json:"is_spicy"`
	Description string    `json:"description"`
	Composition []byte    `json:"composition"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type Foods []Food
