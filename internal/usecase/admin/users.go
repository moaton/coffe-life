package admin

import (
	"coffe-life/config"
	"coffe-life/internal/dto"
	"coffe-life/internal/entity"
	"coffe-life/internal/interfaces"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type users struct {
	repo     interfaces.Repository
	jwtToken config.JwtToken
}

func newUsers(deps Dependencies) *users {
	return &users{
		repo:     deps.Repository,
		jwtToken: deps.JwtToken,
	}
}

var _ interfaces.UsersUsecase = (*users)(nil)

func (u *users) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error) {
	token, err := u.repo.Admin().Users().Login(u.repo.Conn().WithContext(ctx), req, u.jwtToken)
	if err != nil {
		return nil, fmt.Errorf("failed to login: %w", err)
	}
	return &dto.AuthResponse{
		Token: token,
	}, nil
}

func (u *users) CreateUser(ctx context.Context, req dto.CreateUserRequest) error {
	err := u.repo.Admin().Users().CreateUser(u.repo.Conn().WithContext(ctx), convertCreateUserRequestToEntity(req))
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (u *users) GetUsers(ctx context.Context, req dto.GetUsersRequest) ([]*dto.User, error) {
	req.Validate()
	log.Println(req)
	users, err := u.repo.Admin().Users().GetUsers(u.repo.Conn().WithContext(ctx), convertGetUsersRequestToEntity(req))
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	return convertUsersToDto(users), nil
}

func (u *users) GetUserById(ctx context.Context, id string) (*dto.User, error) {
	user, err := u.repo.Admin().Users().GetUserById(u.repo.Conn().WithContext(ctx), id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id %v: %w", id, err)
	}
	return convertUserToDto(user), nil
}

func (u *users) UpdateUser(ctx context.Context, id string, req dto.UpdateUserRequest) error {
	ID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("failed to parse uuid: %w", err)
	}
	user := entity.User{
		ID:        ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	err = u.repo.Admin().Users().UpdateUser(u.repo.Conn().WithContext(ctx), user)
	if err != nil {
		return fmt.Errorf("failed to update user by id %v: %w", id, err)
	}
	return nil
}
