package admin

import (
	"coffe-life/internal/entity"
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

func (r *repository) GetCategories(db *gorm.DB) (entity.Categories, error) {
	var categories entity.Categories
	err := db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *repository) CreateCategory(db *gorm.DB, category entity.Category) (string, error) {
	err := db.Create(&category).Error
	if err != nil {
		return "", err
	}

	return category.ID.String(), nil
}

func (r *repository) UpdateCategory(db *gorm.DB, category entity.Category) error {
	err := db.Save(&category).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteCategory(db *gorm.DB, id string) error {
	err := db.Where("id=?", id).Delete(&entity.Category{}).Error
	if err != nil {
		return err
	}

	return nil
}
