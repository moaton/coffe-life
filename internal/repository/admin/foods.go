package admin

import (
	"coffe-life/internal/entity"
	"coffe-life/internal/interfaces"

	"gorm.io/gorm"
)

type foods struct{}

func newFoods() *foods {
	return &foods{}
}

var _ interfaces.Foods = (*foods)(nil)

func (r *foods) GetFoods(db *gorm.DB) (entity.Foods, error) {
	var foods entity.Foods
	err := db.Find(&foods).Error
	if err != nil {
		return nil, err
	}

	return foods, nil
}

func (r *foods) CreateFood(db *gorm.DB, food *entity.Food) (string, error) {
	err := db.Create(&food).Error
	if err != nil {
		return "", err
	}

	return food.ID.String(), nil
}

func (r *foods) UpdateFood(db *gorm.DB, food *entity.Food) error {
	err := db.Save(&food).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *foods) DeleteFood(db *gorm.DB, id string) error {
	err := db.Where("id=?", id).Delete(&entity.Food{}).Error
	if err != nil {
		return err
	}

	return nil
}
