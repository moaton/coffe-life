package interfaces

import (
	"coffe-life/pkg/gorm/postgres"
)

type Repository interface {
	Conn() *postgres.Gorm
	Admin() AdminRepository
}

type AdminRepository interface {
	Conn() *postgres.Gorm
}
