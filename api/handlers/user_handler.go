package handlers

import (
	"703room/703room.com/models"
	"703room/703room.com/services"
	"703room/703room.com/utils"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	authService services.AuthService
	userService services.UserService
}

func NewUserHandler(authService services.AuthService, userService services.UserService) *UserHandler {
	return &UserHandler{
		authService: authService,
		userService: userService,
	}
}

type SignupRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// POST /api/auth/signup
func (h *UserHandler) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.GetTime(time.April.String())
		var req SignupRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.Error(c, http.StatusBadRequest, utils.ErrInvalidInput, err.Error())
			return
		}
		fmt.Println(req)

		user := models.User{
			Name:     req.Name,
			Email:    req.Email,
			JoinedAt: time.Now(),
		}
		user.PasswordHash = req.Password
		fmt.Println(user)
		token, err := h.authService.Signup(c, &user)
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, utils.ErrSignupFailed, err.Error())
			return
		}

		utils.Created(c, "Signup successful", gin.H{
			"user": gin.H{
				"user_id": user.UserID,
				"name":    user.Name,
				"email":   user.Email,
			},
			"token": token,
		})
	}
}

// POST /api/auth/login
func (h *UserHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("ENTER LOGIN HANDLER")
		var loginData struct {
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&loginData); err != nil {
			utils.Error(c, http.StatusBadRequest, utils.ErrInvalidInput, err.Error())
			return
		}

		user, token, err := h.authService.Login(c, loginData.Email, loginData.Password)
		fmt.Println(loginData.Email)
		fmt.Println(loginData.Password)
		if err != nil {
			utils.Error(c, http.StatusUnauthorized, utils.ErrLoginFailed, err.Error())
			return
		}

		utils.Success(c, "Login successful", gin.H{
			"user": gin.H{
				"user_id": user.UserID,
				"name":    user.Name,
				"email":   user.Email,
			},
			"token": token,
		})
	}
}

// GET /api/users/me
func (h *UserHandler) GetCurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("ENTER GET CURRENT USER HANDLER")

		// Retrieve email from context (set by middleware)
		email, exists := c.Get("email")
		if !exists {
			log.Println("Email not found in context")
			utils.Error(c, http.StatusUnauthorized, utils.ErrUserNotFound, nil)
			return
		}

		userEmail, ok := email.(string)
		log.Println(userEmail)
		if !ok || userEmail == "" {
			log.Println("Email in context is not a valid string")
			utils.Error(c, http.StatusInternalServerError, utils.ErrInvalidUserType, nil)
			return
		}

		// Fetch user from service
		user, err := h.userService.GetUserByEmail(c, userEmail)
		if err != nil {
			log.Printf("Failed to get user by email: %v", err)
			utils.Error(c, http.StatusNotFound, utils.ErrUserNotFound, nil)
			return
		}

		// Return user data
		utils.Success(c, "User fetched successfully", gin.H{
			"user_id": user.UserID,
			"name":    user.Name,
			"email":   user.Email,
		})
	}
}
