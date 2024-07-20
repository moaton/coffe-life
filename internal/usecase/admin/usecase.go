package admin

import (
	"coffe-life/pkg/gorm/postgres"
)

type Dependencies struct {
	DB *postgres.Gorm
}

type Usecase struct {
}

func New(deps Dependencies) *Usecase {
	return &Usecase{}
}
