package admin

import (
	"coffe-life/internal/domain"
	"coffe-life/internal/interfaces"
	"context"
	"fmt"
)

type Dependencies struct {
	Repository interfaces.Repository
}

type Usecase struct {
	repo interfaces.Repository
}

func New(deps Dependencies) *Usecase {
	return &Usecase{
		repo: deps.Repository,
	}
}

func (u *Usecase) GetCategories(ctx context.Context) (domain.Categories, error) {
	categories, err := u.repo.Admin().GetCategories(u.repo.Conn().WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}
	return categories, nil
}
