package interfaces

import (
	"coffe-life/internal/domain"
	"coffe-life/pkg/gorm/postgres"

	"gorm.io/gorm"
)

type Repository interface {
	Conn() *postgres.Gorm
	Admin() AdminRepository
}

type AdminRepository interface {
	Conn() *postgres.Gorm
	GetCategories(db *gorm.DB) (domain.Categories, error)
}
