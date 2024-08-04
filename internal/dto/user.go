package dto

import "github.com/google/uuid"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password" swaggertype:"string" format:"base64"`
}
type AuthResponse struct {
	Token string `json:"token"`
}

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Username  string `json:"username"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type User struct {
	ID        uuid.UUID `json:"id,omitempty"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	IsFirst   bool      `json:"is_first"`
}

type GetUsersRequest struct {
	Page   uint   `json:"page" form:"page"`
	Size   uint   `json:"size" form:"size"`
	Search string `json:"search" form:"search"`
}

func (p *GetUsersRequest) Validate() {
	if p.Page == 0 {
		p.Page = 1
	}
	if p.Size == 0 {
		p.Size = 10
	}
}
