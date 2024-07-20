package interfaces

import (
	"coffe-life/internal/domain"
	"context"
)

type Usecases interface {
	Admin() AdminUsecase
}

type AdminUsecase interface {
	GetCategories(ctx context.Context) (domain.Categories, error)
}
