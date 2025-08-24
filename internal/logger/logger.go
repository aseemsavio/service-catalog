package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New creates a new zap.Logger with the specified log level
func New(level string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	lvl := zapcore.InfoLevel
	_ = lvl.Set(level)
	cfg.Level = zap.NewAtomicLevelAt(lvl)
	cfg.Encoding = "json"
	return cfg.Build()
}
