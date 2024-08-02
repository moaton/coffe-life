package admin

import (
	"coffe-life/config"
	"coffe-life/internal/dto"
	"coffe-life/internal/interfaces"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Dependencies struct {
	Repository interfaces.Repository
	JwtToken   config.JwtToken
}

type Usecase struct {
	repo     interfaces.Repository
	jwtToken config.JwtToken
}

func New(deps Dependencies) *Usecase {
	return &Usecase{
		repo:     deps.Repository,
		jwtToken: deps.JwtToken,
	}
}

func (u *Usecase) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error) {
	token, err := u.repo.Admin().Login(u.repo.Conn().WithContext(ctx), req, u.jwtToken)
	if err != nil {
		return nil, fmt.Errorf("failed to login: %w", err)
	}
	return &dto.AuthResponse{
		Token: token,
	}, nil
}

func (u *Usecase) CreateUser(ctx context.Context, req dto.CreateUserRequest) error {
	err := u.repo.Admin().CreateUser(u.repo.Conn().WithContext(ctx), convertCreateUserRequestToEntity(req))
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
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

func (u *Usecase) GetFoods(ctx context.Context) (dto.Foods, error) {
	foods, err := u.repo.Admin().GetFoods(u.repo.Conn().WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to get foods: %w", err)
	}
	return convertFoodsToDto(foods), nil
}

func (u *Usecase) CreateFood(ctx context.Context, food dto.FoodRequest) (string, error) {
	req, err := convertFoodRequestToEntity(food)
	if err != nil {
		return "", fmt.Errorf("failed to convert food request: %w", err)
	}
	id, err := u.repo.Admin().CreateFood(u.repo.Conn().WithContext(ctx), req)
	if err != nil {
		return "", fmt.Errorf("failed to create food: %w", err)
	}
	return id, nil
}

func (u *Usecase) UpdateFood(ctx context.Context, id string, food dto.FoodRequest) error {
	req, err := convertFoodRequestToEntity(food)
	if err != nil {
		return fmt.Errorf("failed to convert food request: %w", err)
	}
	err = u.repo.Admin().UpdateFood(u.repo.Conn().WithContext(ctx), req)
	if err != nil {
		return fmt.Errorf("failed to update food: %w", err)
	}
	return nil
}

func (u *Usecase) DeleteFood(ctx context.Context, id string) error {
	err := u.repo.Admin().DeleteFood(u.repo.Conn().WithContext(ctx), id)
	if err != nil {
		return fmt.Errorf("failed to delete category %v: %w", id, err)
	}
	return nil
}
