package middlewares

import (
	"703room/703room.com/services"
	"703room/703room.com/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware(authService services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			utils.Error(c, http.StatusUnauthorized, utils.ErrTokenRequired, nil)
			c.Abort()
			return
		}

		user, err := authService.ValidateToken(token)
		if err != nil {
			utils.Error(c, http.StatusUnauthorized, utils.ErrTokenInvalidOrExpire, err.Error())
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
