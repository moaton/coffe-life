package admin

import (
	"coffe-life/internal/domain"
	"coffe-life/pkg/gorm/postgres"

	"gorm.io/gorm"
)

type repository struct {
	db *postgres.Gorm
}

func New(db *postgres.Gorm) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Conn() *postgres.Gorm {
	return r.db
}

func (r *repository) GetCategories(db *gorm.DB) (domain.Categories, error) {
	var categories domain.Categories
	err := db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}
