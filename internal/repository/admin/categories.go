package admin

import (
	"coffe-life/internal/entity"
	"coffe-life/internal/interfaces"

	"gorm.io/gorm"
)

type categories struct{}

func newCategories() *categories {
	return &categories{}
}

var _ interfaces.Categories = (*categories)(nil)

func (r *categories) GetCategories(db *gorm.DB) (entity.Categories, error) {
	var categories entity.Categories
	err := db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *categories) CreateCategory(db *gorm.DB, category *entity.Category) (string, error) {
	err := db.Create(&category).Error
	if err != nil {
		return "", err
	}

	return category.ID.String(), nil
}

func (r *categories) UpdateCategory(db *gorm.DB, category *entity.Category) error {
	err := db.Save(&category).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *categories) DeleteCategory(db *gorm.DB, id string) error {
	err := db.Where("id=?", id).Delete(&entity.Category{}).Error
	if err != nil {
		return err
	}

	return nil
}
