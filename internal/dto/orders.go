package dto

type Order struct {
	ID          string `json:"id"`
	Fullname    string `json:"fullname"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

type Orders []Order
