// config/logger.go
package config

import (
	"log"

	"go.uber.org/zap"
)

func SetupLogger() *zap.Logger {
	logger, err := zap.NewProduction() // Or zap.NewDevelopment() for local
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	return logger
}
