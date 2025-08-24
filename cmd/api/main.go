package main

import (
	"log"
	"net/http"
	"services-catalog/internal/config"
	"services-catalog/internal/logger"
	"services-catalog/internal/migrations"
	"services-catalog/internal/repo"
	"services-catalog/internal/service"

	"go.uber.org/zap"

	httpx "services-catalog/internal/http"
)

func main() {
	configuration := config.Load()
	logg, err := logger.New(configuration.LogLevel)
	if err != nil {
		log.Fatalf("logger: %v", err)
	}
	defer logg.Sync()

	rep, err := repo.Open(configuration.PostgresConnectionString())
	if err != nil {
		logg.Fatal("DB open failed", zap.Error(err))
	}

	err = migrations.RunMigrations(configuration)
	if err != nil {
		logg.Fatal("DB Migrations failed", zap.Error(err))
	} else {
		logg.Info("DB Migrations ran successfully")
	}

	svc := service.New(rep)
	h := httpx.NewHandler(svc)
	r := httpx.NewRouter(h)

	addr := ":" + configuration.HTTPPort
	logg.Info("starting http server", zap.String("addr", addr))
	srv := &http.Server{Addr: addr, Handler: r}
	if err := srv.ListenAndServe(); err != nil {
		logg.Fatal("server stopped", zap.Error(err))
	}
}
