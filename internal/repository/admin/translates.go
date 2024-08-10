package admin

import (
	"coffe-life/internal/entity"
	"coffe-life/internal/interfaces"

	"gorm.io/gorm"
)

type translates struct{}

func newTranslates() *translates {
	return &translates{}
}

var _ interfaces.Translates = (*translates)(nil)

func (r *translates) GetTranslates(db *gorm.DB) ([]entity.Translate, error) {
	var translates []entity.Translate

	err := db.Find(&translates).Error
	if err != nil {
		return nil, err
	}

	return translates, nil
}

func (r *translates) CreateTranslate(db *gorm.DB, translate entity.Translate) error {
	return db.Create(&translate).Error
}

func (r *translates) UpdateTranslate(db *gorm.DB, translate entity.Translate) error {
	return db.Save(&translate).Error
}
