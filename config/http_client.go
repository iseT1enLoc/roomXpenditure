// config/http_client.go
package config

import (
	"net/http"
	"time"
)

func SetupHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 15 * time.Second,
	}
}
