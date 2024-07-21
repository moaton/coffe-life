package entity

import (
	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID `json:"id" gorm:"default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type Categories []Category
