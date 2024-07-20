package main

import (
	"coffe-life/config"
	"coffe-life/internal/application"
	"coffe-life/internal/controller/http"
	"coffe-life/migrations"
	httpserver "coffe-life/pkg/server/http"
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		log.Fatal("failed to init config err:", err)
	}

	err = migrations.RunMigrations(cfg.Postgres.DSN())
	if err != nil {
		log.Fatalf("error running migrations: %s", err)
	}

	app := application.NewWithContext(ctx, cfg)

	app.InitLogger()

	usecases := app.InitUsecases()

	httpRouter := http.NewRouter(http.Dependencies{
		Cfg:      cfg,
		Logger:   app.GetLogger(),
		Usecases: usecases,
	})

	httpServer := httpserver.New(httpRouter, httpserver.Port(cfg.HTTP.Server.Port))

	app.RegisterHTTPServer(httpServer)

	app.Run()
}
