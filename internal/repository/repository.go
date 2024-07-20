package repository

import (
	"coffe-life/internal/interfaces"
	"coffe-life/internal/repository/admin"
	"coffe-life/pkg/gorm/postgres"
)

type repository struct {
	db    *postgres.Gorm
	admin interfaces.AdminRepository
}

func New(db *postgres.Gorm) *repository {
	return &repository{
		db:    db,
		admin: admin.New(db),
	}
}

func (r *repository) Conn() *postgres.Gorm {
	return r.db
}

func (r *repository) Admin() interfaces.AdminRepository {
	return r.admin
}
