package postgres

import (
	"coffe-life/config"
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm struct {
	*gorm.DB
}

func Init(ctx context.Context, conf config.Postgres) (*Gorm, error) {
	conn, err := gorm.Open(postgres.Open(conf.DSN()), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres gorm, err: %w", err)
	}

	db, err := conn.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect postgres gorm, err: %w", err)
	}

	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping postgres gorm, err: %w", err)
	}

	return &Gorm{conn}, nil
}

func (g *Gorm) TxBegin(ctx context.Context) *Gorm {
	return &Gorm{g.WithContext(ctx).Begin()}
}
