package admin

import (
	"coffe-life/internal/dto"
	"coffe-life/internal/interfaces"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type foods struct {
	repo interfaces.Repository
}

func newFoods(deps Dependencies) *foods {
	return &foods{
		repo: deps.Repository,
	}
}

var _ interfaces.FoodsUsecase = (*foods)(nil)

func (u *foods) GetFoods(ctx context.Context) (dto.Foods, error) {
	foods, err := u.repo.Admin().Foods().GetFoods(u.repo.Conn().WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to get foods: %w", err)
	}
	return convertFoodsToDto(foods), nil
}

func (u *foods) CreateFood(ctx context.Context, food dto.Food) (string, error) {
	req, err := convertFoodToEntity(food)
	if err != nil {
		return "", fmt.Errorf("failed to convert food request: %w", err)
	}
	id, err := u.repo.Admin().Foods().CreateFood(u.repo.Conn().WithContext(ctx), req)
	if err != nil {
		return "", fmt.Errorf("failed to create food: %w", err)
	}
	return id, nil
}

func (u *foods) UpdateFood(ctx context.Context, id string, food dto.Food) error {
	req, err := convertFoodToEntity(food)
	if err != nil {
		return fmt.Errorf("failed to convert food request: %w", err)
	}

	ID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("failed to parse uuid: %w", err)
	}

	req.ID = ID
	err = u.repo.Admin().Foods().UpdateFood(u.repo.Conn().WithContext(ctx), req)
	if err != nil {
		return fmt.Errorf("failed to update food: %w", err)
	}
	return nil
}

func (u *foods) DeleteFood(ctx context.Context, id string) error {
	err := u.repo.Admin().Foods().DeleteFood(u.repo.Conn().WithContext(ctx), id)
	if err != nil {
		return fmt.Errorf("failed to delete category %v: %w", id, err)
	}
	return nil
}
