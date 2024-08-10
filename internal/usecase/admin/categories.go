package admin

import (
	"coffe-life/internal/dto"
	"coffe-life/internal/interfaces"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type categories struct {
	repo interfaces.Repository
}

func newCategories(deps Dependencies) *categories {
	return &categories{
		repo: deps.Repository,
	}
}

var _ interfaces.CategoriesUsecase = (*categories)(nil)

func (u *categories) GetCategories(ctx context.Context) (dto.Categories, error) {
	categories, err := u.repo.Admin().Categories().GetCategories(u.repo.Conn().WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}
	return convertCategoriesToDto(categories), nil
}

func (u *categories) CreateCategory(ctx context.Context, category dto.Category) (string, error) {
	id, err := u.repo.Admin().Categories().CreateCategory(u.repo.Conn().WithContext(ctx), convertCategoryToEntity(category))
	if err != nil {
		return "", fmt.Errorf("failed to create category: %w", err)
	}
	return id, nil
}

func (u *categories) UpdateCategory(ctx context.Context, id uuid.UUID, category dto.Category) error {
	req := convertCategoryToEntity(category)
	req.ID = id
	err := u.repo.Admin().Categories().UpdateCategory(u.repo.Conn().WithContext(ctx), req)
	if err != nil {
		return fmt.Errorf("failed to delete category %v: %w", id, err)
	}
	return nil
}

func (u *categories) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	err := u.repo.Admin().Categories().DeleteCategory(u.repo.Conn().WithContext(ctx), id)
	if err != nil {
		return fmt.Errorf("failed to delete category %v: %w", id, err)
	}
	return nil
}
