package admin

import (
	"coffe-life/config"
	"coffe-life/internal/dto"
	"coffe-life/internal/entity"
	"coffe-life/internal/interfaces"
	"coffe-life/internal/utils"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type users struct{}

func newUsers() *users {
	return &users{}
}

var _ interfaces.Users = (*users)(nil)

func (r *users) Login(db *gorm.DB, req dto.LoginRequest, jwt config.JwtToken) (string, error) {
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

func (r *users) CreateUser(db *gorm.DB, user *entity.User) error {
	hash, errGenPass := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if errGenPass != nil {
		return errGenPass
	}
	user.Password = string(hash)
	return db.Create(&user).Error
}

func (r *users) UpdateUser(db *gorm.DB, user entity.User) error {
	return db.Omit("password", "is_first", "username").Save(&user).Error
}

func (r *users) GetUsers(db *gorm.DB, req entity.GetUsersRequest) ([]*entity.User, error) {
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

func (r *users) GetUserById(db *gorm.DB, id string) (*entity.User, error) {
	var user entity.User

	err := db.First(&user, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
