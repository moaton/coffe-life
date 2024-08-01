package admin

import (
	"coffe-life/config"
	"coffe-life/internal/dto"
	"coffe-life/internal/entity"
	"coffe-life/internal/utils"
	"coffe-life/pkg/gorm/postgres"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
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

func (r *repository) Login(db *gorm.DB, req dto.LoginRequest, jwt config.JwtToken) (string, error) {
	var user entity.User
	err := db.Where("username=?", req.Username).First(&user).Error
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", fmt.Errorf("failed to compare hash and password: %w", err)
	}
	tokenString, err := utils.GenerateJwtToken(user, jwt)
	if err != nil {
		return "", fmt.Errorf("failed to generate jwt token: %w", err)
	}
	return tokenString, nil
}

func (r *repository) GetCategories(db *gorm.DB) (entity.Categories, error) {
	var categories entity.Categories
	err := db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *repository) CreateCategory(db *gorm.DB, category *entity.Category) (string, error) {
	err := db.Create(&category).Error
	if err != nil {
		return "", err
	}

	return category.ID.String(), nil
}

func (r *repository) UpdateCategory(db *gorm.DB, category *entity.Category) error {
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

func (r *repository) GetFoods(db *gorm.DB) (entity.Foods, error) {
	var foods entity.Foods
	err := db.Find(&foods).Error
	if err != nil {
		return nil, err
	}

	return foods, nil
}

func (r *repository) CreateFood(db *gorm.DB, food *entity.Food) (string, error) {
	log.Printf("req %s", food.Composition)
	err := db.Debug().Create(&food).Error
	if err != nil {
		return "", err
	}

	return food.ID.String(), nil
}

func (r *repository) UpdateFood(db *gorm.DB, food *entity.Food) error {
	err := db.Save(&food).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteFood(db *gorm.DB, id string) error {
	err := db.Where("id=?", id).Delete(&entity.Food{}).Error
	if err != nil {
		return err
	}

	return nil
}
