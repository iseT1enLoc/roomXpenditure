package middlewares

import (
	serviceimpl "703room/703room.com/services/service_impl"
	"703room/703room.com/utils"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			utils.Error(c, http.StatusUnauthorized, utils.ErrTokenRequired, nil)
			c.Abort()
			return
		}
		data := strings.Split(token, " ")
		if len(data) != 2 {
			c.Abort()
			return
		}
		email, uid, err := serviceimpl.ValidateToken(data[1])
		if err != nil {
			utils.Error(c, http.StatusUnauthorized, utils.ErrTokenInvalidOrExpire, err.Error())
			c.Abort()
			return
		}
		log.Println(*email)
		fmt.Println(*uid)

		c.Set("email", *email)
		c.Set("user_id", *uid)
		c.Next()
	}
}
