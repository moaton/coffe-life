package admin

import (
	"coffe-life/internal/interfaces"
	"coffe-life/pkg/gorm/postgres"
)

type repository struct {
	db         *postgres.Gorm
	users      interfaces.Users
	categories interfaces.Categories
	foods      interfaces.Foods
	translates interfaces.Translates
}

func New(db *postgres.Gorm) *repository {
	return &repository{
		db:         db,
		users:      newUsers(),
		categories: newCategories(),
		foods:      newFoods(),
		translates: newTranslates(),
	}
}

func (r *repository) Conn() *postgres.Gorm {
	return r.db
}

func (r *repository) Users() interfaces.Users {
	return r.users
}

func (r *repository) Categories() interfaces.Categories {
	return r.categories
}

func (r *repository) Foods() interfaces.Foods {
	return r.foods
}

func (r *repository) Translates() interfaces.Translates {
	return r.translates
}
