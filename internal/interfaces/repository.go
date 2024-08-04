package interfaces

import (
	"coffe-life/config"
	"coffe-life/internal/dto"
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

	Login(db *gorm.DB, req dto.LoginRequest, jwt config.JwtToken) (string, error)
	CreateUser(db *gorm.DB, req *entity.User) error
	GetUsers(db *gorm.DB, req entity.GetUsersRequest) ([]*entity.User, error)
	GetUserById(db *gorm.DB, id string) (*entity.User, error)
	UpdateUser(db *gorm.DB, user entity.User) error

	GetCategories(db *gorm.DB) (entity.Categories, error)
	CreateCategory(db *gorm.DB, category *entity.Category) (string, error)
	UpdateCategory(db *gorm.DB, category *entity.Category) error
	DeleteCategory(db *gorm.DB, id string) error

	GetFoods(db *gorm.DB) (entity.Foods, error)
	CreateFood(db *gorm.DB, food *entity.Food) (string, error)
	UpdateFood(db *gorm.DB, food *entity.Food) error
	DeleteFood(db *gorm.DB, id string) error
}
