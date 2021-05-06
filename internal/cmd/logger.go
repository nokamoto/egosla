package cmd

import "go.uber.org/zap"

// NewLogger creates a zap.Logger with a new development or production configuration.
func NewLogger(dev bool) *zap.Logger {
	var cfg zap.Config
	if dev {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
