package interfaces

import (
	"coffe-life/config"
	"context"
)

type Gorm interface {
	Init(ctx context.Context, conf config.Postgres) error
	TxBegin(ctx context.Context) *Gorm
}
