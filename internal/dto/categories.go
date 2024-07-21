package dto

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Categories []Category

type CategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type IdPathParams struct {
	ID string `uri:"id" binding:"required"`
}
