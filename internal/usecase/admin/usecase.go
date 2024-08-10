package admin

import (
	"coffe-life/config"
	"coffe-life/internal/interfaces"
)

type Dependencies struct {
	Repository interfaces.Repository
	JwtToken   config.JwtToken
}

type Usecase struct {
	users      interfaces.UsersUsecase
	categories interfaces.CategoriesUsecase
	foods      interfaces.FoodsUsecase
	translates interfaces.TranslatesUsecase
}

func New(deps Dependencies) *Usecase {
	return &Usecase{
		users:      newUsers(deps),
		categories: newCategories(deps),
		foods:      newFoods(deps),
		translates: newTranslates(deps),
	}
}

func (u *Usecase) Users() interfaces.UsersUsecase {
	return u.users
}

func (u *Usecase) Categories() interfaces.CategoriesUsecase {
	return u.categories
}

func (u *Usecase) Foods() interfaces.FoodsUsecase {
	return u.foods
}

func (u *Usecase) Translates() interfaces.TranslatesUsecase {
	return u.translates
}
