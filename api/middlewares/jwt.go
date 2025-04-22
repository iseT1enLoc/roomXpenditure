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
		fmt.Println(token)
		data := strings.Split(token, " ")
		if len(data) != 2 {
			c.Abort()
			return
		}
		fmt.Println(data[0])
		fmt.Println(data[1])
		email, err := serviceimpl.ValidateToken(data[1])
		if err != nil {
			utils.Error(c, http.StatusUnauthorized, utils.ErrTokenInvalidOrExpire, err.Error())
			c.Abort()
			return
		}
		log.Println(email)

		c.Set("email", *email)
		c.Next()
	}
}
