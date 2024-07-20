package domain

import "github.com/google/uuid"

type Food struct {
	ID           uuid.UUID    `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	Category     string       `json:"category"`
	Price        string       `json:"price"`
	IsNew        bool         `json:"is_new"`
	IsSpicy      bool         `json:"is_spicy"`
	Description  string       `json:"description"`
	Compositions Compositions `json:"compositions"`
}

type Foods []Food
