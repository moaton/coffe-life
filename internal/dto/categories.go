package dto

type Category struct {
	ID          string `json:"id" readonly:"true"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Categories []Category
