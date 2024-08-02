package dto

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
