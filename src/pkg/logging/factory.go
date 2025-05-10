package logging

import (
	"errors"

	"github.com/Hamid-Ba/bama/config"
)

func NewLogger(cfg config.LoggerConfig) (Logger, error) {
	switch cfg.Logger {
	case "zap":
		return NewZapLogger(cfg)
	// case "zerolog":
	//     return NewZerologLogger(cfg)
	default:
		return nil, errors.New("unsupported logger: " + cfg.Logger)
	}
}
