package dto

import "github.com/google/uuid"

type Translate struct {
	ID  uuid.UUID `json:"id" readonly:"true"`
	RU  string    `json:"ru"`
	KZ  string    `json:"kz"`
	ENG string    `json:"eng"`
}
