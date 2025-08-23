package main

import (
	"log"
	"net/http"
	"services-catalog/internal/config"
	"services-catalog/internal/logger"
	"services-catalog/internal/repo"
	"services-catalog/internal/service"

	"go.uber.org/zap"

	httpx "services-catalog/internal/http"
)

func main() {
	cfg := config.Load()
	logr, err := logger.New(cfg.LogLevel)
	if err != nil {
		log.Fatalf("logger: %v", err)
	}
	defer logr.Sync()

	rep, err := repo.Open(cfg.PGDSN())
	if err != nil {
		logr.Fatal("db open failed", zap.Error(err))
	}

	svc := service.New(rep)
	h := httpx.NewHandler(svc)
	r := httpx.NewRouter(h)

	addr := ":" + cfg.HTTPPort
	logr.Info("starting http server", zap.String("addr", addr))
	srv := &http.Server{Addr: addr, Handler: r}
	if err := srv.ListenAndServe(); err != nil {
		logr.Fatal("server stopped", zap.Error(err))
	}
}
