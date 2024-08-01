package entity

import "github.com/google/uuid"

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
}

type Foods []Food
