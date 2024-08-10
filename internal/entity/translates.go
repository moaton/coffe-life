package entity

import "github.com/google/uuid"

type Translate struct {
	ID  uuid.UUID `gorm:"default:uuid_generate_v4()"`
	RU  string
	KZ  string
	ENG string
}
