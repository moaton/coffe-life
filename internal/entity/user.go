package entity

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id" gorm:"default:uuid_generate_v4()"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"password"`
	Username  string    `json:"username"`
	IsFirst   bool      `json:"is_first"`
}

type Pagination struct {
	Limit  int
	Offset int
}

type GetUsersRequest struct {
	Pagination `json:"pagination"`
	Search     string `json:"search"`
}
