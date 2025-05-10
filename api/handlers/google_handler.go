package handlers

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"

	"703room/703room.com/services"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// GoogleHandler holds the Google OAuth service
type GoogleHandler struct {
	googleService services.IGoogleService
	clientID      string
	clientSecret  string
	redirectURL   string
}

// NewGoogleHandler creates a new handler
func NewGoogleHandler(service services.IGoogleService, clientID, clientSecret, redirectURL string) *GoogleHandler {
	return &GoogleHandler{
		googleService: service,
		clientID:      clientID,
		clientSecret:  clientSecret,
		redirectURL:   redirectURL,
	}
}

func (h *GoogleHandler) RedirectToGoogle(c *gin.Context) {
	state := h.googleService.GenerateStateOauthCookie(c.Writer)

	oauthGoogleConfig := &oauth2.Config{
		ClientID:     h.clientID,
		ClientSecret: h.clientSecret,
		RedirectURL:  h.redirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
	}

	url := oauthGoogleConfig.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *GoogleHandler) HandleGoogleCallback(c *gin.Context) {
	log.Println("Enter Google OAuth callback")

	code := c.Query("code")
	state := c.Query("state")

	oauthState, err := c.Cookie("oauthstate")
	if err != nil || state != oauthState {
		log.Println("Invalid OAuth state:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state"})
		return
	}

	oauthGoogleConfig := &oauth2.Config{
		ClientID:     h.clientID,
		ClientSecret: h.clientSecret,
		RedirectURL:  h.redirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
	}

	token, err := oauthGoogleConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Failed to exchange code for token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token exchange failed"})
		return
	}

	client := oauthGoogleConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Println("Failed to get user data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
	defer resp.Body.Close()

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read response body:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read user info"})
		return
	}

	jwtToken, err := h.googleService.GoogleLogin(c.Request.Context(), userData)
	if err != nil {
		log.Println("Token generation failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}
	front_end_url := os.Getenv("FRONT_END_URL")

	redirectURL := front_end_url + "/rooms?token=" + jwtToken
	c.Redirect(http.StatusPermanentRedirect, redirectURL)
}
