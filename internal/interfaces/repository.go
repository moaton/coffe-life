package interfaces

import (
	"coffe-life/internal/entity"
	"coffe-life/pkg/gorm/postgres"

	"gorm.io/gorm"
)

type Repository interface {
	Conn() *postgres.Gorm
	Admin() AdminRepository
}

type AdminRepository interface {
	Conn() *postgres.Gorm
	GetCategories(db *gorm.DB) (entity.Categories, error)
	CreateCategory(db *gorm.DB, category entity.Category) (string, error)
	UpdateCategory(db *gorm.DB, category entity.Category) error
	DeleteCategory(db *gorm.DB, id string) error
}
