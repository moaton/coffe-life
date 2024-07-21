package interfaces

import (
	"coffe-life/internal/dto"
	"context"
)

type Usecases interface {
	Admin() AdminUsecase
}

type AdminUsecase interface {
	GetCategories(ctx context.Context) (dto.Categories, error)
	CreateCategory(ctx context.Context, category dto.CategoryRequest) (string, error)
	UpdateCategory(ctx context.Context, id string, category dto.CategoryRequest) error
	DeleteCategory(ctx context.Context, id string) error
}
