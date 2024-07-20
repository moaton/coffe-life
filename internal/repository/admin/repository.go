package admin

import "coffe-life/pkg/gorm/postgres"

type repository struct {
	db *postgres.Gorm
}

func New(db *postgres.Gorm) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Conn() *postgres.Gorm {
	return r.db
}
