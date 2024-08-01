package dto

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password" swaggertype:"string" format:"base64"`
}
type LoginResponse struct {
	Token string `json:"token"`
}
