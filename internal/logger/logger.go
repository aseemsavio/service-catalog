package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(level string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	lvl := zapcore.InfoLevel
	_ = lvl.Set(level)
	cfg.Level = zap.NewAtomicLevelAt(lvl)
	cfg.Encoding = "json"
	return cfg.Build()
}
