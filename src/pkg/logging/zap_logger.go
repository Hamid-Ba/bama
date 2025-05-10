package logging

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Hamid-Ba/bama/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	zap *zap.Logger
}

func NewZapLogger(cfg config.LoggerConfig) (Logger, error) {
	// Set logging level
	var level zapcore.Level
	if err := level.Set(cfg.Level); err != nil {
		level = zapcore.InfoLevel // fallback
	}

	// Ensure log directory exists
	if err := os.MkdirAll(cfg.FilePath, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	// Use filepath.Join for OS-safe path
	appLogPath := filepath.Join(cfg.FilePath, "app.log")
	// errorLogPath := filepath.Join(cfg.FilePath, "error.log")

	// Setup Zap config
	zapCfg := zap.Config{
		Encoding:    cfg.Encoding,
		Level:       zap.NewAtomicLevelAt(level),
		OutputPaths: []string{appLogPath},
		// ErrorOutputPaths: []string{errorLogPath},
		EncoderConfig: zap.NewProductionEncoderConfig(),
	}

	// Optional: Set time format and caller info
	zapCfg.EncoderConfig.TimeKey = "timestamp"
	zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapCfg.EncoderConfig.CallerKey = "caller"

	// Build logger
	z, err := zapCfg.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build zap logger: %w", err)
	}

	return &ZapLogger{zap: z}, nil
}

func (l *ZapLogger) Info(msg string, fields ...Field) {
	l.zap.Info(msg, convert(fields...)...)
}
func (l *ZapLogger) Debug(msg string, fields ...Field) {
	l.zap.Debug(msg, convert(fields...)...)
}
func (l *ZapLogger) Error(msg string, fields ...Field) {
	l.zap.Error(msg, convert(fields...)...)
}
func (l *ZapLogger) Warn(msg string, fields ...Field) {
	l.zap.Warn(msg, convert(fields...)...)
}
func (l *ZapLogger) Sync() error {
	return l.zap.Sync()
}

func convert(fields ...Field) []zap.Field {
	zfs := make([]zap.Field, len(fields))
	for i, f := range fields {
		zfs[i] = zap.Any(f.Key, f.Value)
	}
	return zfs
}
