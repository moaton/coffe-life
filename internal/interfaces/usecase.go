package interfaces

import (
	"coffe-life/internal/dto"
	"context"
)

type Usecases interface {
	Admin() AdminUsecase
}

type AdminUsecase interface {
	Login(ctx context.Context, category dto.LoginRequest) (*dto.AuthResponse, error)
	CreateUser(ctx context.Context, req dto.CreateUserRequest) error
	GetUsers(ctx context.Context, req dto.GetUsersRequest) ([]*dto.User, error)
	GetUserById(ctx context.Context, id string) (*dto.User, error)
	UpdateUser(ctx context.Context, id string, req dto.UpdateUserRequest) error

	GetCategories(ctx context.Context) (dto.Categories, error)
	CreateCategory(ctx context.Context, category dto.CategoryRequest) (string, error)
	UpdateCategory(ctx context.Context, id string, category dto.CategoryRequest) error
	DeleteCategory(ctx context.Context, id string) error

	GetFoods(ctx context.Context) (dto.Foods, error)
	CreateFood(ctx context.Context, category dto.FoodRequest) (string, error)
	UpdateFood(ctx context.Context, id string, category dto.FoodRequest) error
	DeleteFood(ctx context.Context, id string) error
}
