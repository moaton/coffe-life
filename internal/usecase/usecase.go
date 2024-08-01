package usecase

import (
	"coffe-life/config"
	"coffe-life/internal/interfaces"
	"coffe-life/internal/usecase/admin"
)

type Dependencies struct {
	Repository interfaces.Repository
	JwtToken   config.JwtToken
}

type Usecases struct {
	admin interfaces.AdminUsecase
	repo  interfaces.Repository
}

func New(deps Dependencies) *Usecases {
	adminDeps := admin.Dependencies{
		Repository: deps.Repository,
		JwtToken:   deps.JwtToken,
	}

	return &Usecases{
		admin: admin.New(adminDeps),
		repo:  deps.Repository,
	}
}

func (u *Usecases) Admin() interfaces.AdminUsecase {
	return u.admin
}
