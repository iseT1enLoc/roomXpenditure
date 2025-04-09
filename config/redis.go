// config/redis.go
package config

import (
	"github.com/go-redis/redis/v8"
)

func SetupRedis(addr, password string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,     // e.g., "localhost:6379"
		Password: password, // "" if no password
		DB:       db,       // default is 0
	})
	return rdb
}
