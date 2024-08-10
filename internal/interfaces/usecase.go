package interfaces

import (
	"coffe-life/internal/dto"
	"context"

	"github.com/google/uuid"
)

type Usecases interface {
	Admin() AdminUsecase
}

type AdminUsecase interface {
	Users() UsersUsecase
	Categories() CategoriesUsecase
	Foods() FoodsUsecase
	Translates() TranslatesUsecase
}

type UsersUsecase interface {
	Login(ctx context.Context, category dto.LoginRequest) (*dto.AuthResponse, error)
	CreateUser(ctx context.Context, req dto.CreateUserRequest) error
	GetUsers(ctx context.Context, req dto.GetUsersRequest) ([]*dto.User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*dto.User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, req dto.UpdateUserRequest) error
}

type CategoriesUsecase interface {
	GetCategories(ctx context.Context) (dto.Categories, error)
	CreateCategory(ctx context.Context, category dto.Category) (string, error)
	UpdateCategory(ctx context.Context, id uuid.UUID, category dto.Category) error
	DeleteCategory(ctx context.Context, id uuid.UUID) error
}

type FoodsUsecase interface {
	GetFoods(ctx context.Context) (dto.Foods, error)
	CreateFood(ctx context.Context, category dto.Food) (string, error)
	UpdateFood(ctx context.Context, id uuid.UUID, category dto.Food) error
	DeleteFood(ctx context.Context, id uuid.UUID) error
}

type TranslatesUsecase interface {
	GetTranslates(ctx context.Context) ([]dto.Translate, error)
	CreateTranslate(ctx context.Context, translate dto.Translate) error
	UpdateTranslate(ctx context.Context, translate dto.Translate) error
}
