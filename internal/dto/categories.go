package dto

import "github.com/google/uuid"

type Category struct {
	ID          string `json:"id" readonly:"true"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Categories []Category

type IdPathParams struct {
	ID uuid.UUID `uri:"id" binding:"required"`
}
