package admin

import (
	"coffe-life/internal/dto"
	"coffe-life/internal/interfaces"
	"context"
	"fmt"

	"github.com/google/uuid"
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

func (u *Usecase) GetCategories(ctx context.Context) (dto.Categories, error) {
	categories, err := u.repo.Admin().GetCategories(u.repo.Conn().WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}
	return convertCategoriesToDto(categories), nil
}

func (u *Usecase) CreateCategory(ctx context.Context, category dto.CategoryRequest) (string, error) {
	id, err := u.repo.Admin().CreateCategory(u.repo.Conn().WithContext(ctx), convertCategoryRequestToEntity(category))
	if err != nil {
		return "", fmt.Errorf("failed to create category: %w", err)
	}
	return id, nil
}

func (u *Usecase) UpdateCategory(ctx context.Context, id string, category dto.CategoryRequest) error {
	ID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("failed to parse id %v: %w", id, err)
	}
	req := convertCategoryRequestToEntity(category)
	req.ID = ID
	err = u.repo.Admin().UpdateCategory(u.repo.Conn().WithContext(ctx), req)
	if err != nil {
		return fmt.Errorf("failed to delete category %v: %w", id, err)
	}
	return nil
}

func (u *Usecase) DeleteCategory(ctx context.Context, id string) error {
	err := u.repo.Admin().DeleteCategory(u.repo.Conn().WithContext(ctx), id)
	if err != nil {
		return fmt.Errorf("failed to delete category %v: %w", id, err)
	}
	return nil
}
