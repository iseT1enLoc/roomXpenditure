package services

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

type IGoogleService interface {
	GoogleLogin(ctx context.Context, data []byte) (string, error)
	GetUserDataFromGoogle(googleOauthConfig *oauth2.Config, code, oauthGoogleUrlAPI string) ([]byte, error)
	GenerateStateOauthCookie(w http.ResponseWriter) string
}
