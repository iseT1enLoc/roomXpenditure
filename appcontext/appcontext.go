package appcontext

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppContext struct {
	db         *gorm.DB
	redis      *redis.Client
	logger     *zap.Logger
	config     *viper.Viper
	httpClient *http.Client
}

// NewAppContext creates a new context with dependencies
func NewAppContext(
	db *gorm.DB,
	redis *redis.Client,
	logger *zap.Logger,
	httpClient *http.Client,
) *AppContext {
	return &AppContext{
		db:         db,
		redis:      redis,
		logger:     logger,
		httpClient: httpClient,
	}
}

// Getters for dependency injection
func (a *AppContext) DB() *gorm.DB {
	return a.db
}

func (a *AppContext) Redis() *redis.Client {
	return a.redis
}

func (a *AppContext) Logger() *zap.Logger {
	return a.logger
}

func (a *AppContext) Config() *viper.Viper {
	return a.config
}

func (a *AppContext) HTTPClient() *http.Client {
	return a.httpClient
}
