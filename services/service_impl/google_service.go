package serviceimpl

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"log"

	"703room/703room.com/models"
	"703room/703room.com/repository"
	"703room/703room.com/services"

	"github.com/google/uuid"

	"gorm.io/gorm"

	"golang.org/x/oauth2"
)

type googleService struct {
	userRepo    repository.UserRepository
	authService services.AuthService
}

func NewGoogleService(userRepo repository.UserRepository, authService services.AuthService) services.IGoogleService {
	return &googleService{
		userRepo:    userRepo,
		authService: authService,
	}
}

func (lu *googleService) GenerateStateOauthCookie(w http.ResponseWriter) string {
	expiration := time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}
func (g *googleService) GetUserDataFromGoogle(googleOauthConfig *oauth2.Config, code string, oauthGoogleUrlAPI string) ([]byte, error) {
	// Exchange the authorization code for a token
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code for token: %v", err)
	}

	// Use the token to create an authenticated client
	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get(oauthGoogleUrlAPI)
	if err != nil {
		return nil, fmt.Errorf("failed to get user data from Google: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get user data: status %d, body: %s", resp.StatusCode, body)
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read user data response: %v", err)
	}

	return userData, nil
}

func (g *googleService) GoogleLogin(ctx context.Context, data []byte) (string, error) {
	var googleUser models.GoogleUser
	err := json.Unmarshal(data, &googleUser)
	if err != nil {
		log.Println("Error unmarshalling Google user data:", err)
		return "", err
	}

	// Construct user model from Google data
	user := &models.User{
		UserID:   uuid.New(),
		GoogleId: googleUser.Id,
		Email:    googleUser.Email,
		Name:     googleUser.Name,
	}

	// Check if user already exists
	existingUser, err := g.userRepo.GetByEmail(ctx, googleUser.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Error checking existing user:", err)
		return "", err
	}

	// If user does not exist, insert into DB
	if existingUser == nil {
		if err := g.userRepo.CreateUser(ctx, user); err != nil {
			log.Println("Error inserting new user:", err)
			return "", err
		}
	} else {
		user.UserID = existingUser.UserID // maintain ID if found
	}

	// Generate token
	token, err := g.authService.GenerateToken(user)
	if err != nil {
		log.Println("Error generating token:", err)
		return "", err
	}

	return token, nil
}
