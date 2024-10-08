package application

import (
	"coffe-life/config"
	"coffe-life/internal/interfaces"
	"coffe-life/internal/repository"
	"coffe-life/internal/usecase"
	"coffe-life/pkg/gorm/postgres"
	"coffe-life/pkg/logger/zap"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-logr/logr"
	"go.uber.org/zap/zapcore"
)

type Application struct {
	ctx        context.Context
	cfg        *config.Config
	db         *postgres.Gorm
	logger     logr.Logger
	httpServer interfaces.Server

	shutdown chan os.Signal
}

func NewWithContext(ctx context.Context, cfg *config.Config) (*Application, error) {
	db, err := postgres.Init(ctx, cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("failed to init gorm: %w", err)
	}

	app := &Application{
		ctx:      ctx,
		cfg:      cfg,
		db:       db,
		shutdown: make(chan os.Signal, 1),
	}

	return app, nil
}

func (a *Application) Run() {
	a.logger.Info("Application started...")

	<-a.shutdown
	a.Stop()
}

func (a *Application) Stop() {
	close(a.shutdown)

	a.logger.Info("Application stoped")
}

func (a *Application) InitUsecases() interfaces.Usecases {
	deps := usecase.Dependencies{
		Repository: repository.New(a.db),
		JwtToken:   a.cfg.JwtToken,
	}

	return usecase.New(deps)
}

func (a *Application) InitLogger() {
	a.logger = zap.New(
		zap.Level(zapcore.DebugLevel),
		zap.UseDevMode(true),
		zap.TimeEncoder(zapcore.ISO8601TimeEncoder),
		zap.ConsoleEncoder(
			func(ec *zapcore.EncoderConfig) { ec.EncodeLevel = zapcore.CapitalColorLevelEncoder },
			func(ec *zapcore.EncoderConfig) {
				ec.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
			},
			func(ec *zapcore.EncoderConfig) {
				ec.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
					encoder.AppendString(filepath.Base(caller.FullPath()))
				}
			},
		),
	)
}

func (a *Application) GetLogger() logr.Logger {
	return a.logger
}

func (a *Application) RegisterHTTPServer(server interfaces.Server) {
	a.httpServer = server
}
