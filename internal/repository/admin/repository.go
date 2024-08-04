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

func (r *repository) CreateUser(db *gorm.DB, user *entity.User) error {
	return db.Create(&user).Error
}

func (r *repository) UpdateUser(db *gorm.DB, user entity.User) error {
	return db.Omit("password", "is_first", "username").Save(&user).Error
}

func (r *repository) GetUsers(db *gorm.DB, req entity.GetUsersRequest) ([]*entity.User, error) {
	var users []*entity.User
	q := db

	if req.Search != "" {
		s := "%" + req.Search + "%"
		q = q.Where("LOWER(username) like ?", s).
			Or("LOWER(first_name) like ?", s).
			Or("LOWER(last_name) like ?", s)
	}

	err := q.Debug().Limit(req.Limit).Offset(req.Offset).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) GetUserById(db *gorm.DB, id string) (*entity.User, error) {
	var user entity.User

	err := db.First(&user, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
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
