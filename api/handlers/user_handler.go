package handlers

import (
	"703room/703room.com/models"
	"703room/703room.com/services"
	"703room/703room.com/utils"
	"fmt"
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
		fmt.Println(user)
		// Hash password before storing (you can do this in service or here)
		hashed, err := h.authService.HashPassword(req.Password)
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, "Password hashing failed", err.Error())
			return
		}
		user.PasswordHash = hashed
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
func (h *UserHandler) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		utils.Error(c, http.StatusBadRequest, utils.ErrInvalidInput, err.Error())
		return
	}

	user, token, err := h.authService.Login(c.Request.Context(), loginData.Email, loginData.Password)
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

// GET /api/users/me
func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, utils.ErrUserNotFound, nil)
		return
	}

	u, ok := user.(*models.User)
	if !ok {
		utils.Error(c, http.StatusInternalServerError, utils.ErrInvalidUserType, nil)
		return
	}

	utils.Success(c, "User fetched successfully", gin.H{
		"user_id": u.UserID,
		"name":    u.Name,
		"email":   u.Email,
	})
}
